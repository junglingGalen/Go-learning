package main

import (
	"fmt"
	"sort"
	"time"
	"math/rand"
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

type Hero struct {
	Name string
	Age int
}

type Heros []Hero

func (h Heros) Len()int  {
	return len(h)
}

func (h Heros) Less(i, j int) bool {
	if h[i].Age < h[j].Age {
		return true
	}
	return false
}

func (h Heros) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
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

	//构建英雄切片
	var heros Heros
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 5; i++ {
		time.Sleep(1*time.Nanosecond)
		hero := Hero {
			Name : fmt.Sprintf("%v", time.Now().UnixNano()),
			Age : rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	fmt.Println("当前英雄列表=", heros)

	sort.Sort(heros)
	fmt.Println("英雄切片按年龄排序后=", heros)
}