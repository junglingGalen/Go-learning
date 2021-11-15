package main
import (
	"fmt"
)

/*数组的初步使用*/
func array1()  {
	// 定义一个数组
	var hans [6]float64
	// 给数组中的每一个元素复制
	hans[0] = 3.0
	hans[1] = 5.0
	hans[2] = 1.0
	hans[3] = 3.4
	hans[4] = 2.0
	hans[5] = 50.0
	// 遍历数组求和
	totalWeight2 := 0.0
	for i := 0; i < len(hans); i++ {
		totalWeight2 += hans[i]
	}
	// 求出平均体重
	avgWeight2 := fmt.Sprintf("%.2f", totalWeight2 / 6)
	fmt.Printf("taotalWeight=%v avgWeight2=%v", totalWeight2, avgWeight2)
}

/*数组详解1*/
func array2() {
	var intArr [3]int  // 64位电脑int占8个字节
	// 当我们定义完数组之后，其实数组的各个元素都有默认值

	intArr[0], intArr[1], intArr[2] = 1, 2, 3

	fmt.Println(intArr)
	// 数组的地址可以通过数组名获取&intArr，就是数组的首元素的地址，数组的第二个元素是上一个元素的地址+元素所占空间
	fmt.Println("数组的地址：", &intArr, "第一个元素的地址：", &intArr[0], "第二个元素的地址：", &intArr[1])

}

// func main() {
// 	array2()
// }

/*练习：从终端输入5个成绩，保存到float64数组，并输出。*/
func array3() {
	var score [5]float64
	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d个元素的值\n", i+1)
		fmt.Scanln(&score[i])
	}
	fmt.Println(score)
}

// func main() {
	// array3()
// }

/*数组初始化的四种方式*/
func array04() {
	var numArr01 [3]int = [3]int{1, 2, 3}
	var numArr02 = [3]int{1, 2, 3}
	var numArr03 = [...]int{8, 9, 10}
	var numArr04 = [...]int{1: 100, 0: 900, 3:253}  // 未填充的下标为默认值
	fmt.Println(numArr01, numArr02, numArr03, numArr04)
}

// func main() {
// 	array04()
// }

/*数组遍历的几种方法*/
func array05() {
	var numArr01 = [3]int{1, 2, 3}
	for index, value := range numArr01 {
		fmt.Println(index, value)
	}
	for i := 0; i < len(numArr01); i++ {
		fmt.Println(i, numArr01[i])
	}
}

func main() {
	array05()
}