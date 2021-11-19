package main

import (
	"fmt"
)

/*
为什么需要切片：保存数量不固定的数据
切片的基本介绍： 
①切片是数组的一个引用，因此切片是引用数据类型，在进行传递时，
遵守应用的传递机制
② 切片的使用和数组类似，遍历切片、访问切片的元素和求切片的长度都一样
③ 切片的长度是可以变化的，因此切片是一个可以动态变化的数组
④ 切片定义的基本语法：
	var 变量名 []类型
*/

func slice() {
	// 基本使用
	var intArr [5]int = [...]int{1,22,33,44,55}
	// 声明一个切片
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是=", slice)  // 22, 23
	fmt.Println("slice 的元素个数是=", len(slice))  // 2
	fmt.Println("slice 的容量是=", cap(slice))  // 4  切片的容量是随着切片长度的增长自增长的

	// 切片是引用，地址看对应元素和数组相同
	fmt.Println("slice的0号元素的地址是=", &slice[0])  
	fmt.Println("intArr的1号元素地址是=", &intArr[1])  
	// 切片其实就是一个数据结构，它的地址里只记录了自己第一个元素的地址、自身长度、容量


}

// func main() {
// 	slice()
// }

/*切片的定义方式*/
func sliceCreate() {
	// 1.切数组获得切片
	// 2.make([]type, len, [cap])  cap > len
	var slice []float64 = make([]float64, 5, 6)
	slice[2] = 3
	fmt.Println(slice)
	// 3.定义一个切片，直接就指定具体数组，使用原理类似make方式
	var slice2 []string = []string{"Tom", "jsck", "mary"}
	fmt.Println("strSlice=", slice2, len(slice2), cap(slice2))
}

func main()  {
	sliceCreate()
}

