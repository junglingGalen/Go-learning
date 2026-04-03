package main

import (
	"fmt"
	"13.factoryMode/model"
)

func main() {
	var stu = model.NewStudent("tom", 18.8)
	fmt.Println(stu.Name)
	fmt.Println(stu.Score())
}