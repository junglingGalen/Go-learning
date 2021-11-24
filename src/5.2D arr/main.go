package main

import (
	"fmt"
)

func printArr() {
	/*
	请用二维数组输出
	000000
	001000
	020300
	000000
	*/
	var _arr [4][6] int  //初始化默认全部零值
	var arr [4][6]int = [4][6]int{[6]int{0,0,0,0,0,0}, [6]int{0,0,1,0,0,0}, [6]int{0,2,0,3,0,0}, [6]int{0,0,0,0,0,0}}
	// 遍历打印出图形
	for i := 0; i < len(arr); i++ {
		for _, j := range arr {
			fmt.Printf("%v", i)
		}
		fmt.Println('\n')
	}
}

/*多维数组，外层数组保存的是内层数组的地址，只有最内层数组存的是值*/

func main() {
	printArr()
}