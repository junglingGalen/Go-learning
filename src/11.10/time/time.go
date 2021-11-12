package main

import (
	"time"
	"fmt"
	"strconv"
)

func test03() {
	str := ""
	for i := 0; i < 200000; i++ {
		str += "hello" + strconv.Itoa(i)
		if i % 50000 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Println("最终耗时:", end - start)
	time.Sleep(2 * time.Minute)
}