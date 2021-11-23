package main

import (
	"fmt"
	"encoding/json"
)

/* 结构体使用细节 */

func myStruct1() {
	/* 
	1.结构体的值类型数据在内存中是连续存储的；指针类型数据指针地址是连续，但指向的地址不一定是连续的

	*/
	type Point struct {
		x int
		y int
	}

	type Rect struct {
		leftUp, rightDown Point
	}

	r := Rect{Point{15, 20}, Point{20, 99}}
	fmt.Println(r)
	fmt.Println(&r.leftUp.x, &r.leftUp.y, &r.rightDown.x, &r.rightDown.y)
	// 0xc000010400 0xc000010408 0xc000010410 0xc000010418 

	type Rect2 struct {
		leftUp, rightDown *Point
	}

	r2 := Rect2{&Point{15, 20}, &Point{20, 99}}
	fmt.Println(r2)
	fmt.Println(&r2.leftUp, &r2.rightDown)  // 连续的0xc00004a1f0 0xc00004a1f8
	fmt.Printf("指向的地址是%p, %p\n", r2.leftUp, r2.rightDown)
	fmt.Println(&r2.leftUp.x, &r2.leftUp.y, &r2.rightDown.x, &r2.rightDown.y)
	// 0xc0000120c0 0xc0000120c8 0xc0000120d0 0xc0000120d8
}

// func main() {
// 	myStruct1()
// }

func myStruct2() {
	/*结构体之间的转换*/

	/* 1.结构体是用户单独定义的类型，只有字段完全一致（名字、类型、数量）才可以相互转换 */
	type A struct {
		Num int
	}

	type B struct {
		Num int
	}

	var a A
	var b B
	a = A(b)  // √ 这里是允许的，字段完全一致（名字、类型、数量）可以转换

	/*2.结构体进行type重新定义，Golang人为是新的数据类型，但是相互间可以强转*/
	type Student struct {
		Name string
		Age int
	}

	type Student2 Student

	var stu1 Student
	var stu2 Student2

	// stu2 = stu1  // 错误，不能直接复制，重定义后是两种数据类型，只有 强转后才可以赋值
	stu2 = Student2(stu1)
	fmt.Println(stu2, a)

}

func myStuct3() {
	/*struct的每个字段上，可以写一个tag，
	该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化*/
	type Monster struct{
		Name string `json:"name"`
		Age int `json:"age"`
		Skill string `json:"skill"`
	}
	monster := Monster{"牛魔王", 500, "牛头拳"}

	jsonStr, err := json.Marshal(monster)  // monster小写字段无法在json包里运用，小写只能本包作用
	if err != nil {
		fmt.Println("json 处理错误", err)
	}
	fmt.Println(string(jsonStr))

}

func main() {
	myStuct3()
}
