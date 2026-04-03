package main

import (
	"fmt"
)

/*编写结构体MethodUtils，编程一个方法，方法不需要参数，在方法中打印10*8矩阵*/
type MethodUtils struct {

}

func (util *MethodUtils) print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%v", "*")
		}
		fmt.Printf("%v", "\n")
	}
}

func (util *MethodUtils) print2(m int, n int)(s int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%v", "*")
		}
		fmt.Printf("%v", "\n")
	}
	s = m * n
	return s
}

func main() {
	m := MethodUtils{}
	(&m).print()
	fmt.Println((&m).print2(6, 9))
}