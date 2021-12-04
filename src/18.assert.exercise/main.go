package main

import (
	"fmt"
)

type I struct{}

func TypeJudge(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {  // 这里type是关键字，固定的写法
			case bool:
				fmt.Printf("param #%d is a bool 值是%v\n", i, x)
			case float32:
				fmt.Printf("param #%d is a float32 值是%v\n", i, x)
			case float64, int:
				fmt.Printf("param #%d is a int/float64 值是%v\n", i, x)
			case nil:
				fmt.Printf("param #%d is a nil 值是%v\n", i, x)
			case string:
				fmt.Printf("param #%d is a string 值是%v\n", i, x)
			case *I:
				fmt.Printf("param #%d is a *I 值是%v\n", i, x)
			default:
				fmt.Printf("param #%d is a unknow 值是%v\n", i, x)
		}
	}
}

func main() {
	var k interface{}
	var i I
	TypeJudge(true, 3.2, 4, float32(2.8), k, i, &i)
}