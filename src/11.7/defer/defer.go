package main

import (
	"fmt"
)

/*
当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈中
当函数执行完毕后，再从defer栈，按照先入后出的方式出栈执行
defer将语句放入栈时，也会将相关的值深拷贝同时入栈。
*/

func sum(n1 int, n2 int) int {
	
	defer fmt.Println("ok n1=", n1) //defer
	defer fmt.Println("ok n2=", n2) //defer
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}