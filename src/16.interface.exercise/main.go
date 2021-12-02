package main

import (
	"fmt"
	"sort"
)

/* 最佳实践
	1.	实现对Hero结构体切片的排序：sort.Sort(data interface)
*/

func main() {
	var intSlice = []int{0, -1, 10, 7, 90}
	/* 要求对intSlice进行排序 */
	// 1. 写个冒泡排序函数
	// 2. 使用内置的排序包
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// 请大家对结构体切片进行排序

}