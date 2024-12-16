package main

import "fmt"

type usb interface {
	start()
	stop()
}

type phone struct {
	Name string
}

type ca struct {
	Name string
}

type computer struct{}

func (c computer) test(u usb) {
	u.start()
	u.stop()
}

func (c ca) start() {
	fmt.Println(c.Name, "start")
}
func (c ca) stop() {
	fmt.Println(c.Name, "stop")
}

func (ph phone) start() {
	fmt.Println(ph.Name, "start")
}

func (ph phone) stop() {
	fmt.Println(ph.Name, "stop")
}

func main() {
	var iphone phone
	iphone.Name = "apple"
	var u usb = iphone
	u.start()

	var computers computer
	computers.test(iphone)
}
