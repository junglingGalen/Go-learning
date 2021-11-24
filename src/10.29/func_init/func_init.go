package main
import (
	"fmt"
)

var b = c
var t = test()
var d = t + b

var c int

func test() int {
	fmt.Println("test已运行")
	fmt.Printf("b的值时%v", b)
	return 5
}

func init() {
	fmt.Println("init已运行", d)
}

func main() {
	fmt.Println("main已运行")
}  