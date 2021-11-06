package main
 
 import (
	 "fmt"
	 _ "time"
 )

 func main() {
	 var money float64= 100000
	 var times int
	 label:
	 for {
		 if money > 50000 {
			 money -= money * 0.05
			 times++
		 } else {
			 if money >= 1000 {
				 money -= 1000
				 times++
			 } else {
				 break label
			 }
		 }
	 }
	 fmt.Printf("总共循环了%v次", times)
	 
	 fmt.Println("程序结束")
 }