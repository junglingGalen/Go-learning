package main

import (
	"fmt"
	"14.OopFeature/model"
)

/* 面向对象变成三大特性

一、封装
	1.封装的概念
	封装就是把抽象的字段和对字段的操作封装在一起，数据被保护在内部，程序的其它包只能通过被授权
	的操作（方法），才能对字段进行操作
	2.go里实现封装
	① 将结构体、字段的首字母小写
	② 给结构体提供一个工厂模式的函数，首字母大写。类似一个构造函数
	③ 提供一个首字母大写的Set方法对属性进行修改
	④ 提供一个首字母大写的Get方法获取属性的值

二、继承
	1.继承的概念
	当多个结构体存在相同的属性和方法时，可以从这些结构体中抽象出结构体，在该结构体中定义这些相同的属性和方法
	其它结构体不需要重新定义这些属性和方法，只需要嵌套一个抽象出来的该结构体即可

	也就是说：Golang中，如果一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法
	从而实现继承

	2.golang继承的特性
	① 结构体可以使用嵌套结构体的所有属性和方法（包括小写首字母属性）
	② 机构体的属性查找，先从自身属性查找，找不到会去自己继承的匿名结构体查找（所以pupil.Student.Score可以写成pupil.Score）
	如果都找不到则报错
	③ 当结构体和匿名机构体有相同字段或者方法时，编译器采用就近访问原则访问，如果希望访问
	匿名结构体的字段和方法，可以通过匿名结构体名来区分
	④ 多重继承： 结构体嵌入两个（或多个）匿名结构体，如两个匿名结构体有相同的字段和方法（同时结构体本身没有同名字段和方法）
	在访问时，就必须明确指定匿名结构体类型，否则编译器报错（建议不适用多重继承）
	⑤ 如果一个struct嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在访问
	组合的结构体字段或方法时，必须带上结构体的名字
	⑥ 继承的结构体和当前机构体有同名方法时，访问匿名结构体的方法需要先访问匿名结构体
	⑦ 匿名结构体也可以是基本数据类型（int,string...等等）,同一个数据类型的匿名字段只能有一个
*/

type Student struct {
	Name string
	Score float64
}

func (stu *Student) test() {
	fmt.Printf("%v is testing\n", stu.Name)
}

type Pupil struct {
	Student  // 嵌套匿名结构体
	Toy string
}

type Graduate struct {
	stu Student  // 嵌套有名结构体，该模式是组合
	GirlFriend string
}

func (p *Pupil) test() {
	fmt.Println("小学生在考试。。")
}

func main() {
	/* 1. 封装的实际使用 */
	per := model.NewPerson("jack", 12, 5000.5)
	fmt.Println("姓名：", per.Name, "年龄：", per.GetAge(), "薪水：", per.GetSalary())

	/* 2. 继承的实际使用 */
	
	// 匿名结构体的创建方式1
	pupil := Pupil{}
	pupil.Name = "Lajue" // 等价于pupil.Student.Name
	pupil.Student.Score = 98.5
	pupil.Toy = "plane"

	// 匿名结构体的创建方式2
	pupil2 := Pupil{Student{"Tom", 12.3}, "Monster"}  // 嵌套匿名结构体，也可以在创建结构体实例时，直接指定各个匿名结构体字段的值


	pupil.Student.test() // 继承的结构体和当前机构体有同名方法时，访问匿名结构体的方法需要先访问匿名结构体
	pupil.test()

	pupil2.Student.test()

}