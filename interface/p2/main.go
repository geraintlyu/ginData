package main

import "fmt"

type use interface {
	start()
	stop()
}

type phone struct {
	Name string
}

func (ph *phone) start() {
	fmt.Println(ph.Name, "start")
}
func (ph *phone) stop() {
	fmt.Println(ph.Name, "stop")
}

func main() {
	var iphone phone = phone{
		Name: "apple",
	}

	/*
		由于上面的start和stop两个方法使用的是指针接收者
		所以下面这种写法是错误的
		var u use = iphone
		u.start()
		u.stop()
		# command-line-arguments
		./main.go:25:14: cannot use iphone (variable of type phone) as use value in
		variable declaration: phone does not implement use (method start has pointer
		receiver)
	*/
	var u use = &iphone // 对于方法是指针接受者，需要将地址复制给接口变量
	u.start()
	u.stop()
}
