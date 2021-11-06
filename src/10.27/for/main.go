package main

import (
	"fmt"
)

func main() {
	// 传统的遍历方法是按照字节遍历的，
	// 一次取出一个字节，而go是按utf-8编码的中文占3个字节,所以中卫乱码
	// 转化成切片才可以
	var str string
	str = "hello world你好啊"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}

	// 第二种遍历方式
	str = "abc-ok上海"
	for index, val := range str {
		fmt.Printf("index=%d, value=%c \n", index, val)
	}
}