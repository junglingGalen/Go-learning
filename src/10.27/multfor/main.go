package main

import (
	"fmt"
)


// func main()  {
// 	// 1)统计3个班成绩情况，每个班有5名学生，
// 	// 求出各个班的平均分和所有班级的平均分[学生的成绩从键盘输入]
// 	var sumSocre float64
// 	for i := 1; i <= 3; i++ {
// 		var classScore float64
// 		for j := 1; j <= 5; j++ {
// 			var score float64
// 			fmt.Println("请输入同学成绩：")
// 			fmt.Scanln(&score)
// 			classScore += score
// 		}
// 		fmt.Printf("你们班的平均成绩是%v\n", classScore / 5.0)
// 		sumSocre += classScore
// 	}
// 	fmt.Printf("\n所有同学的平均成绩是%v\n", sumSocre / 15.0)
// }

func main(){
	// 打印金字塔经典案例
	// 请编写一个程序，可以接收一个整数表示层数，打印出金字塔
	var tar int
	fmt.Println("请输入金字塔层数：")
	fmt.Scanln(&tar)
	for i := 1; i <= tar; i++ {
		for j := 1; j <= 2 * tar - 1; j++ {
			if j >= tar - i && j < tar + i -1 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}