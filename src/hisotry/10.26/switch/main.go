package main

import (
	"fmt"
)

func main() {
	var i byte
	fmt.Println("请输入字母：")
	fmt.Scanf("%c", &i)
	switch i {
	case 'a':
		fmt.Println("周一，猴子穿新衣")
		fallthrough
	case 'b':
		fmt.Println("周二，猴子当小二")
	default:
		fmt.Println("猴子真帅")
	}
}