package main

import (
	"fmt"
)

/* 方法介绍
1. 结构体是值类型，在方法调用中遵守值类型的传递机制，是拷贝传递方式
2. 入程序员希望在方法中，修改结构体变量的值，可以通过结构体指针的方式来处理
3. Golang中的方法作用在指定的数据类型上的（即：和指定的数据类型绑定），因此自定义类型
都可以有方法，而不仅仅是struct，比如int，float32等都可以有方法
4. 方法的访问范围控制的规则，和函数一样，方法名首字母小写，只能在本包访问，方法首字母大写，
可以在本包和其它访问。
5. 如果一个变量实现了String()这个方法，那么fmt.Println默认会调用这个变量的String()进行输出
*/

type Person struct{
	Name string
}

//给A类型绑定一份方法
func (p Person) test() {  // p为当前结构体
	fmt.Println("test()", p.Name)
	p.Name = "Jack"  // 方法传参和函数传参一样，结构体是值传参，不会修改掉原结构体
}

// func main() {
// 	var p Person
// 	p.Name = "Tom"
// 	p.test()
// 	fmt.Println(p.Name)  // 方法传参和函数传参一样，结构体是值传参，不会修改掉原结构体

// 	/* 方法的调用和传参机制
// 	方法的调用和传参机制和函数基本一至，不一样的地方是调用时会将调用方法的变量当做实参
// 	传给方法 */

// }

type Circle struct {
	radius float64
}

func (circle Circle) Perimeter() (perimeter float64) {
	perimeter = 3.14 * 2 * circle.radius
	return perimeter
}

func (circle *Circle) Perimeter2() (perimeter float64) {
	/* 为了高效我们传入指针，可以更便捷的修改原元素，结构体go底层做了处理，支持直接指针.属性 */
	circle.radius = 10
	return circle.Perimeter()
}

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

func (i *integer) change() {
	*i = *i + 1
}

func main() {
	c := Circle{5}
	fmt.Println(c.Perimeter(), (&c).Perimeter2(), c.Perimeter())
	var i integer = 1
	(&i).change()
	i.print()
}