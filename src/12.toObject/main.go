package main

import (
	"fmt"
)

/*
学生案例：
编写一个student结构体，包含name、gender、age、id、score字段，分别为string、string、int、int、float
结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段
在main方法中，创建Student结构体实例，并访问say方法，并将调用结果打印输出
*/

type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]\n",
		student.name, student.gender, student.age, student.id, student.score,  // 这里不加)的话需要加，
	)
	return infoStr
}

func main() {
	var stu = Student{"Tom", "male", 18, 1000, 99.98}
	var stu2 = Student{
		name:"Tom", 
		gender:"male", 
		age:18, 
		id:1000, 
		score:99.98}
	fmt.Println(stu.say(), stu2.say())
}