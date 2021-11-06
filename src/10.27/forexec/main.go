package main

import (
	"fmt"
)


func main()  {

	// 打印1~100之间所有9的倍数的证书的个数及总和
	var sum int
	for i := 1; i <= 100; i++ {
		if i % 9 == 0{
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Printf("总和为%d\n", sum)

	// 完成下面表达式的输出
	var y interface{}
	var x int
	fmt.Scanln(&x)
	y = x
	switch y.(type) {
		case int:
			for i := 0; i <= x; i++ {
				fmt.Printf("%d + %d = %d\n", i, x - i, x)
			}
		default:
			fmt.Println("请输入正整数！")
	}
	


}
