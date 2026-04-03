package main
import (
	"fmt"
	"math"
)

func main() {
	// exec1
	var n1 int32 = 10
	var n2 int32 = 50
	if n1 > n2 {
		fmt.Println("hello world")
	} else {
		fmt.Println("你好啊")
	}

	// exec2
	fmt.Println("\nexec2")
	var n3 float64 = 11.0
	var n4 float64 = 17.0
	if n3 > 10.0 && n4 < 20 {
		fmt.Println("和=", n3 + n4)
	}

	// exec 3
	fmt.Println("\nexec3")
	var n5 int32 = 12
	var n6 int32 = 17
	if tmp1, tmp2 := (n5 + n6) % 3, (n5 + n6) % 5;  tmp1 == 0 && tmp2 == 0 {
		fmt.Println("都可以整除")
	} else {
		fmt.Printf("n5 = %d, n6 = %d, n5 + n6除3余%d,除5余%d\n", n5, n6, tmp1, tmp2)
	}

	// exec4
	fmt.Println("\nexec4")
	var year int
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		fmt.Println("是闰年~")
	} else {
		fmt.Println("不是闰年！")
	}

	// exec5
	fmt.Println("\nexec5")
	var a float64 = 3.0
	var b float64 = 100.0
	var c float64 = 60.0
	m := 2 * b - 4 * a * c
	if m > 0 {
		x1 := (-b + math.Sqrt(2 * b - 4 * a *c)) / (2 * a)
		x2 := (-b - math.Sqrt(2 * b - 4 * a *c)) / (2 * a)
		fmt.Printf("该算数式有解%v和%v", x1, x2)
	} else if m == 0 {
		fmt.Printf("该算数式有解%v", -b/(2 * a))
	} else {
		fmt.Println("该函数无解")
	}
}

