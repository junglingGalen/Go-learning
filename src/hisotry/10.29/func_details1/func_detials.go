package main

import (
	"fmt"
)

// 可变参数
func sumAll(num... int) (sum int) {
	for _, i := range num {
		sum += i
	}
	return
}

func main() {
	fmt.Println(sumAll(1,2,3,4,5,6,7))
}