package main

import (
	"fmt"
	"time"
	"math/rand"
)

/*
（1）创建一个byte类型的26个元素的数组，分别放置'A'-'Z'。使用for循环访问所有元素
并打印出来。提示：字符数据运算'A' + 1 -> 'B'
*/
func excise1() {
	var arr = [26]byte{}
	var i byte = 'A'
	for ; i <= 'Z'; i++ {
		arr[i - 'A'] = i
	}
	for _, i := range arr {
		fmt.Printf("%c", i)
	}
}

// func main() {
// 	excise1()
// }

/*
(2)请求出一个数组的最大值并得到对应的下标
*/
func getMax(arr [25]int) (max int, index int) {
	for idx, value := range arr {
		if index == 0 {
			index, max = idx, value
		} else {
			if value > max {
				max, index = value, idx
			}
		}
	}
	return max, index
}

// func main() {
// 	fmt.Println(getMax([25]int{1:15, 24:45, 13:89}))
// }

/*
(3)要求：随机生成5个数，并将其反转打印。
*/
func randomArr() {
	// rand.Seed(time.Now().UnixNano() + int64(1))
	time.Now().UnixNano()
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn(100) + 1
	}
	for i := 4; i >= 0; {
		fmt.Println(arr[i])
		i -= 1
	}
}

func main() {
	randomArr()
}

/*
数组的复杂运用
*/
// 1.翻转