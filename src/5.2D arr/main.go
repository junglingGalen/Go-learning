package main

import (
	"fmt"
)

func main() {
	/*
	请用二维数组输出
	000000
	001000
	020300
	000000
	*/
	var arr [4][6]int = [4][6]int{[6]int{0,0,0,0,0,0}, [6]int{0,0,1,0,0,0}, [6]int{0,2,0,3,0,0}, [6]int{0,0,0,0,0,0}}
	// 遍历打印出图形
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v", i)
		}
		fmt.Println('\n')
	}
}