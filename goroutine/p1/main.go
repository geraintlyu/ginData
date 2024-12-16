package main

import (
	"fmt"
	"time"
)

/*
	golang通过协程来实现多线程，由于协程的消耗非常小，单独创建一个协程可能只需要消耗2kb
	类似于Java创建一个线程可能需要消耗2mb的内存
	协程是用户级的，线程是内核级的
*/

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test(), 运行-", i)
		time.Sleep(time.Microsecond * 100)
	}
}

/*
golang通过go关键字来启动一个新的协程
*/
func main() {
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main(), 运行-", i)
		time.Sleep(time.Microsecond * 100)
	}

	fmt.Println("主进程退出...")
}
