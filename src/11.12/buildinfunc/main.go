package main

import (
	"fmt"
)

func main() {
	num1 := 100
	fmt.Printf("num1的类型%T，num1的值%v, num1的地址%v\n", num1, num1, &num1)

	// new用来分配内存主要给值分配内存
	num2 := new(int) // *int
	fmt.Printf("num2的类型%T，num2的值%v, num2的地址%v, num2的值时%v\n", num2, num2, &num2, *num2)

	// make用来分配内存，主要用来分配引用类型
	
}