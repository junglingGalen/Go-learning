package main

import (
	"fmt"
	"math"
)

// func main() {

// 	// 判断一个整数属于哪个范围
// 	var inp int
// 	fmt.Println("请输入一个整数:")
// 	fmt.Scanln(&inp)
// 	switch {
// 	case inp < 0:
// 		fmt.Println("小于0")
// 	case inp == 0:
// 		fmt.Println("等于0")
// 	case inp > 0:
// 		fmt.Println("大于0")
// 	}

// }

func main() {
	// 判断一个整数是否是水仙花数
	var inp int
 	fmt.Println("请输入一个整数:")
	fmt.Scanln(&inp)
	hundred := inp / 100
	ten := (inp - hundred * 100) / 10
	one := inp - hundred * 100 - ten * 10
	sum :=  math.Pow(float64(hundred), 3) + math.Pow(float64(ten), 3)
	sum += math.Pow(float64(one), 3)
	if sum == float64(inp) {
		fmt.Println("是水仙花数")
	} else {
		fmt.Println("不是水仙花数")
	}
}