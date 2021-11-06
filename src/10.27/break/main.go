 package main
 
 import (
	 "fmt"
	 "math/rand"
	 "time"
 )

 func main() {
	 var i int
	 for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println(n)
		i++
		if n == 99{
			break
		}
	}
	fmt.Printf("随机生成了%v次产生了99", i)
 }