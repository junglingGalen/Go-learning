package main

import (
	"fmt"
	"sort"
)

/* 最佳实践
	1.	实现对Hero结构体切片的排序：sort.Sort(data interface)
*/

type MySlice struct {
	s []int
}

func (m MySlice) Len() int {
	return len(m.s)
}

func (m MySlice) Less(i, j int) bool {
	if m.s[i] < m.s[j] {
		return true
	} else {
		return false
	}
}

func (m MySlice) Swap(i, j int) {
	temp := m.s[i]
	m.s[i] = m.s[j]
	m.s[j] = temp
}

func main() {
	var intSlice = []int{0, -1, 10, 7, 90}
	/* 要求对intSlice进行排序 */
	// 1. 写个冒泡排序函数
	// 2. 使用内置的排序包
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// 请大家对结构体切片进行排序
	var myslice = MySlice{[]int{1,5,8,4,1,2,3,6,4,9,8}}
	sort.Sort(myslice)
	fmt.Println(myslice.s)
}