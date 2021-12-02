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
	fmt.Printf("%T\n", a)
	var point Point = Point{1, 2}
	fmt.Printf("%T\n", point)
	a = point  // OK空接口可以接受任何类型的实例
	fmt.Printf("%T\n", a)  // main.Point。打印出来的是具体类型不是接口类型
	var b Point
	// b = a  // cannot use a (type interface {}) as type Point in assignment: need type assertion
	b = a.(Point)  // 类型断言，接口实例转换成其它类型时需要断言
	/* 表示判断a是否指向Point类型的变量，
	如果是就转成Point类型并赋值给b变量，否则报错 */
	fmt.Println(b)
}