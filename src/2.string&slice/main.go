package main

import (
	"fmt"
)

func main() {
	//string底层是一个byte数组，因此string也是可以进行切片处理
	str := "hello@hhw"
	//使用切片获取到@hhw
	slice := str[5:]
	fmt.Println("slice=", slice)

	//string是不可变的，也就说不能通过str[0] = 'z'方式来修改字符串
	//str[0] = 'z' [编译器不通过，报错，原因string是不可变的]

	//修改字符串中的某个值
	//string -> []byte / 或者[]rune -> 修改 -> 重写成string

	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println(str)

	//细节，我们转成[]byte后，可以处理应为和数字，但是不能处理中文，原因是[]byte是按字节来处理，而一个汉字是3个字节
	//解决方法是，将string转成[]rune即可，因为[]rune是按字符处理，兼容汉字

	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println("str=", str)
}