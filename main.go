package main

import (
	"fmt"
)

func sum(x ...int) int {
	var cal int = 0
	for i := 0; i < len(x); i++ {
		cal += x[i]
	}
	return cal
}

func main() {
	slice := [3]int{1, 2, 3}
	fmt.Printf("slice: %T\n", slice)
	var arr1 = [...]int{1, 2, 3}
	fmt.Printf("arr1: %T\n", arr1)
	var slice1 = make([]int, 4, 5)
	fmt.Printf("slice1: %v\n", slice1)
}
