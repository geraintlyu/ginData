package main

import (
	"fmt"
	"reflect"
)

/*
golang的反射机制能够处理空接口
1. reflect.TypeOf()获取变量的类型
2. reflect.ValueOf()获取变量的值
3. 如果修改空接口中变量的值，需要通过set[TYPE]()函数来进行修改

修改空接口中变量的值
1. 如果要修改存储在空接口的变量的值，首先需要确保空接口中存储的是变量的指针
2. 在修改变量之前，需要通过Kind()来获取变量的详细类型，然后才能通过set[TYPE]来修改
*/
func fu1(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	if t.Elem().Kind() == reflect.Int {
		v.Elem().SetInt(20)
	}
}

func main() {
	a := 10
	fu1(&a)
	fmt.Println(a)
}
