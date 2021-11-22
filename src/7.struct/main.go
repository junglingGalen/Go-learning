package main

import (
	"fmt"
)

func myStruct1() {
	/* 
	结构体简介
	① 结构体是自定义的数据类型，代表一类事物
	② 结构体变量（实例）是具体的，实际的，代表一个具体变量
	③ 结构体变量是值类型
	*/

	/* 如何声明结构体 */
	type Cat struct {
		Name string
		color string
		age int
	}

	/* 
	字段的注意事项和使用细节
	① 字段声明语法同变量，示例：字段名 字段类型
	② 字段的类型可以为： 基本类型、数组或引用类型
	③ 在创建一个结构体变量后，如果没有给字段赋值，都对应一个零值，规则同前面讲的一样
		bool类型为false，数值类型为0，string类型是""
		数组类型的默认值和它的元素类型有关，比如score[3]int 则为[0,0,0]
		指针，slice，和map的零值都是nil，即还没有分配空间
	④ ！！！不同结构体变量之间的字段是独立的，一个变量修改不影响其他变量，结构体默认为值拷贝
	*/
	type Person struct {
		Name string
		Age int
		Scores [5]float64
		ptr *int
		slice []int
		map1 map[string]string
	}
	var person1 Person
	fmt.Println(person1)  // { 0 [0 0 0 0 0] <nil> [] map[]} 其实几个引用类型都是nil只是输出不同而已

	person1.slice = make([]int, 10)
	person1.map1 = make(map[string]string)
	fmt.Println(person1)

	// 和数组一样结构体默认是值拷贝
	arr := [5]int{1,2,3,4,5}
	arr2 := arr
	arr2[0] = 100
	fmt.Println(arr, arr2)

	var cat1 Cat
	cat1.Name = "baobao"
	cat1.age = 12
	cat1.color = "black"

	cat2 := cat1
	cat2.color = "white"

	fmt.Println(cat1, cat2)
}

// func main() {
// 	myStruct1()
// }

func myStruct2() {
	/* 结构体实例化的几种方式 */
	
	type Person struct{
		Name string
		Age int
	}

	// ① var p Person 声明后再复制
	var p1 Person

	// ② p2 := Person{}
	p2 := Person{}

	// ③ var person *Person = new(Person)
	var p3 *Person = new(Person)
	(*p3).Name = "smith"
	p3.Age = 30
	/* 因为p3是一个指针，因此标准的字段赋值方式是
	(*p3).Name = "smith"  也可以这样写 p3.Age = 30
	原因：go的设计者为了程序员方便，底层会对p3.Name = "smith"进行处理，会给p3加上取值运算 (*p3).Name = "smith"
	*/

	// ④ var person *Person = &Person{}
	var p4 *Person = &Person{}

	fmt.Println(p1, p2, *p3, *p4)
}

func main() {
	myStruct2()
}