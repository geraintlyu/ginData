package main

import (
	"fmt"
	"sync"
	"time"
)

// 对于读取锁来说，读取可以并发读取不受限制
func lock1() {
	defer wg.Done()
	mutex.RLock()
	fmt.Println("---执行读取操作")
	time.Sleep(time.Second * 2)
	mutex.RUnlock()
}

// 对于写入锁来说，同一时间只允许一个协程进行写入
func lock2() {
	defer wg.Done()
	mutex.Lock()
	fmt.Println("执行写入操作")
	time.Sleep(time.Second * 2)
	mutex.Unlock()
}

var wg sync.WaitGroup

// var mutex sync.Mutex // 可以通过这个语句来定义一个锁
var mutex sync.RWMutex // 可以通过这个语句来定义一个读写锁

func main() {
	for i := 0; i < 7; i++ {
		wg.Add(1)
		go lock1()
	}
	for i := 0; i < 7; i++ {
		wg.Add(1)
		go lock2()
	}
	wg.Wait()
}
