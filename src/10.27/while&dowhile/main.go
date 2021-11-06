package main

import (
	"fmt"
)

func main() {
	var i int = 1
	for ;;i++{
		if i > 10 {
			break
		}
		fmt.Println(i)
	}
}