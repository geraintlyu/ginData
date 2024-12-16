package main

import "fmt"

func main() {
	// channel是一个引用类型，定义的时候需要初始化
	ch := make(chan int, 3)

	// 将数据写入管道，管道遵循FIFO
	ch <- 10
	ch <- 20
	ch <- 30

	// 从管道内部获取数据,并且将数据赋值给变量
	cha := <-ch
	fmt.Printf("cha: %v\n", cha)

	// 从管道内部取出数据，但是不保留。虽然没有保留，但是管道内不还是会丢失一个数据
	<-ch

	/*
		循环遍历一个管道, 遍历管道有两种方式
		1. for-range 这种方式遍历管道要求管道在遍历之前处于关闭状态，如果遍历之前没有关闭，会导致deadlock
		2. for 这种方式遍历之前不要求管道处于关闭状态，即使管道没有关闭也能够遍历
	*/

	ch1 := make(chan int, 4)
	ch2 := make(chan int, 4)

	for i := 1; i < 4; i++ {
		ch1 <- i
		ch2 <- i
	}

	for i := 1; i < 4; i++ {
		fmt.Println(<-ch2)
	}

	// 下面的例子中，for-range如果在遍历之前没有关闭ch1，则会引发deadlock
	close(ch1)
	for v := range ch1 {
		fmt.Printf("v: %v\n", v)
	}
}
