package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

/* 类型断言
由于接口是一般类型，不知道具体类型，吐过要转成具体类型，就需要使用类型断言
*/

func main() {
	var a interface{}
	fmt.Printf("%T\n", a)  // nil

	var point Point = Point{1, 2}
	fmt.Printf("%T\n", point)  // main.Point

	a = point  // OK空接口可以接受任何类型的实例
	fmt.Printf("%T\n", a)  // main.Point。打印出来的是具体类型不是接口类型
	 
	var b Point
	// b = a  // cannot use a (type interface {}) as type Point in assignment: need type assertion
	b = a.(Point)  // 类型断言，接口实例转换成其它类型时需要断言

	/* 表示判断a是否指向Point类型的变量，
	如果是就转成Point类型并赋值给b变量，否则报错 */

	fmt.Println(b)


	/* 类型断言的其它案例 */
	var x interface{}
	var b2 float32 = 1.1
	x = b2 // 空接口，可以接受任意类型

	y := x.(float32)
	fmt.Printf("y的类型是%T 值是%v\n", y, y)

	y2, _ := x.(float64)  // _ 接收错误信息，不会报错程序不会挂
	fmt.Printf("y2的类型是%T 值是%v\n", y2, y2)  // y2的类型是float64 值是0(报错但也会赋值)

	if y3, succ := x.(float64); succ {    // 
		fmt.Printf("y3的类型是%T 值是%v\n", y3, y3)
	} else {
		fmt.Println("类型转换失败！")
	}
	fmt.Println(y3)

}