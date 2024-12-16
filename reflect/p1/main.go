package main

import (
	"fmt"
	"reflect"
)

type Persion struct {
	Name    string `json:"username"`
	Age     int    `json:"age"`
	Account int    `json:"account"`
}

func (p Persion) GetPersion() {
	fmt.Println(p.Name)
}

func (p *Persion) SetPersion(age int) {
	p.Age = age
}

// reflect能够动态的操作结构体，包括动态的查看结构体的内容：包括结构体的字段的类型、字段Tag、字段的名字等；
// 还可以动态地调用结构体的方法

// 需要注意在修改结构体的内容的时候，需要像操作空接口一样，向函数传入指向结构体的指针
func reflectPersion(p1 interface{}) {
	t := reflect.TypeOf(p1)
	v := reflect.ValueOf(p1)

	// 确保递交给函数的是一个Persion类型的变量
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println(p1, "is not struct")
	}

	// 获取结构体中所有的字段的类型
	for i := 0; i < t.Elem().NumField(); i++ {
		fmt.Println(t.Elem().Field(i))
	}

	// 获取结构体中所有的字段的值
	for i := 0; i < v.Elem().NumField(); i++ {
		fmt.Println(v.Elem().Field(i))
	}

	if methods := v.Elem().MethodByName("GetPersion"); methods.IsValid() {
		methods.Call([]reflect.Value{})
	} else {
		fmt.Println("methods not found")
	}

	// 对于引用传递的SetPersion(age int)函数来说，这里不需要使用v.Elem()，直接通过v.MethodByName()来调用结构体的函数
	// 但是对于值传递的GetPersion()函数来说，在调用函数的时候使用v.Elen()或者直接v.MethodByName()都是可以的
	if methods := v.MethodByName("SetPersion"); methods.IsValid() {
		methods.Call([]reflect.Value{reflect.ValueOf(20)})
	} else {
		fmt.Println("SetPersion is not defind")
	}
}

func main() {
	var p1 Persion = Persion{
		Name:    "geraint",
		Age:     18,
		Account: 100026004,
	}

	reflectPersion(&p1)
	fmt.Printf("p1: %#v\n", p1)
}
