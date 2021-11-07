package main

import (
	"fmt"
)

func main() {
	//1.统计字符串的长度（按字节数返回）
	str := "start北"
	fmt.Printf("%T-%v\n", len(str), len(str))

	// 字符串遍历，同时处理有中文的问题
	str2 := "hello北京"
	fmt.Printf("%T\n", str2)
	for i:=0; i < len(str2); i++ {
		fmt.Printf("字符=%x\n,type=%T", str2[i], str2[i])
	}
}