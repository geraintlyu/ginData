package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. 单项管道：只读管道 <-chan int；只写管道 chan<- int
// 2. select多路复用
// 3. goroutine panic处理错误

func test() {
	defer wg.Done()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err happend: %v\n", err)
		}
	}()
	var data map[int]string
	data[0] = "test"
}

func f2() {
	defer wg.Done()
	fmt.Println("hello world!")
}

// testChan1()函数的形参是一个只读的管道，这个管道只能够支持读取数据，如果试图写入数据会触发错误
func testChan1(ch <-chan int) {
	defer wg.Done()
	for value := range ch {
		fmt.Printf("value: %v\n", value)
		time.Sleep(time.Microsecond * 50)
	}
}

// testChan2()函数的形参是一个只写的管道，这个管道仅支持写入，如果试图读取数据会触发错误
func testChan2(ch chan<- int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("向管道内写入数据:", i)
		time.Sleep(time.Microsecond * 50)
	}
	close(ch)
}

var wg sync.WaitGroup

func main() {
	/*
	   此处由于test()这个协程对一个没有指定空间的map进行了赋值，
	   所以test()协程在运行的时候会导致整个程序停止。
	   此时在test()中定义的错误处理机制导致了整个程序依旧能够正常运行
	*/
	wg.Add(1)
	go test()
	wg.Add(1)
	go f2()

	/*
	   可以通过只读管道和只写管道来限制管道的行为，使得管道更加的安全
	*/
	ch := make(chan int, 4)
	wg.Add(1)
	go testChan1(ch)
	wg.Add(1)
	go testChan2(ch)
	wg.Wait()

	var (
		chT1 = make(chan int, 3)
		chT2 = make(chan int, 3)
		chT3 = make(chan int, 3)
	)

	for i := 0; i < 3; i++ {
		chT1 <- i
		chT2 <- i
		chT3 <- i
	}

	/*
	   使用select关键字来实现多路复用，select会选择一个随机的、已经准备好的管道来进行处理
	*/
	select {
	case va := <-chT1:
		fmt.Printf("va: %v\n", va)
	case vb := <-chT2:
		fmt.Printf("vb: %v\n", vb)
	case vc := <-chT2:
		fmt.Printf("vc: %v\n", vc)
	default:
		fmt.Println("管道为空")
	}

}
