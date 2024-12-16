package main

import (
	"fmt"
	"sync"
)

func test1(ch chan int) {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		ch <- i
		fmt.Printf("写入数据:%v\n", i)
		// time.Sleep(time.Microsecond * 50)
	}
	close(ch)
}

func test2(ch chan int) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("读取数据:%v\n", val)
		// time.Sleep(time.Microsecond * 50)
	}
}

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 10)

	wg.Add(1)
	go test1(ch)

	wg.Add(1)
	go test2(ch)

	wg.Wait()
	fmt.Println("退出")
}
