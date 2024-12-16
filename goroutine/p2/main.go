package main

import (
	"fmt"
	"sync"
	"time"
)

/*
需要注意,如果主进程(main进程)的运行速度比协程快，那么默认情况下主进程不会等待其他协程
一旦主进程结束，程序就会立刻推出。

常见的解决方法是创建一个进程计数器，来实现让主进程等待其他子进程
*/

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1(), 运行-", i)
		time.Sleep(time.Microsecond * 100)
	}
	// 每次协程执行完成之后，就让协程计数器 -1
	defer wg.Done()
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2(), 运行-", i)
		time.Sleep(time.Microsecond * 100)
	}

	// 每次协程执行完成之后，就让协程计数器 -1
	defer wg.Done()
}

// 创建一个协程计数器
var wg sync.WaitGroup

func main() {
	//每次开启一个协程就让协程计数器 +1
	wg.Add(1)
	go test1()

	//每次开启一个协程就让协程计数器 +1
	wg.Add(1)
	go test2()

	//让主进程等待所有协程执行完成之后再退出
	wg.Wait()
	fmt.Println("主进程退出...")
}
