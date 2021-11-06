package main
import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("请输入你的年龄：")

	fmt.Scanln(&age)
	if age := 184; age >= 18 {  // go支持在if条件内定义变量，只能作用域后续代码块
		fmt.Printf("%v岁了，你要对自己的行为负责！", age)
	} else {
		fmt.Printf("happy吧，你还小！")
	}

	
}
