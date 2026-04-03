package main

import (
	"fmt"
)

var name string = "Tom"
// name := "Rose"  // 错误，该语句相当于两个语句，一个初始化变量，一个赋值语句属于运算语句，不能再函数外运行

func test() {
	name = "Jason"
}

func main() {
	fmt.Println(name)
	test()
	fmt.Println(name)

	for i := 1; i <=1; i++ {
		fmt.Println(i)
		var j int64 = 255
		fmt.Println(j)
	}
}