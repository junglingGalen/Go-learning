package main

import (
	"fmt"
)

/*
编写一个函数fbn(n int),要求完成
①可以接受一个n int
②能够将斐波那契的数列放到切片中
*/

func fbn(n int) []int{
	var slice []int = make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			slice[i] = 1
		} else {
			slice[i] = slice[i - 1] + slice[i - 2]
		}
	}
	return slice
}

func main() {
	fmt.Println(fbn(10))
}