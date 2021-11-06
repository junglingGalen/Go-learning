package main

import (
	"fmt"
)


// 案例
// func test(n int) {
// 	if n > 2 {
// 		n--
// 		test(n)
// 	}
// 	fmt.Println("n=", n)
// }

// func main() {
// 	test(4)
// }

//斐波那契树，请使用递归的方式，求出斐波那契树1,1,2,3,5,8,13，求第n个数
func febo(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else {
		return febo(n - 1) + febo(n - 2)
	}
}

func main() {
	fmt.Println(febo(4))
}