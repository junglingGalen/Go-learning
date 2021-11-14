package main
import (
	"fmt"
	"math/rand"
	"time"
)

// 循环打印输入的月份的天数
func judgeIsLeapYear(year int) (res bool) {
	if year % 100 != 0 && year % 4 == 0 || year % 100 == 0 && year % 400 == 0 {
		res = true
	} else {
		res = false
	}
	return res
}

func getMonthDayCount(month int, isLeapYear bool) int {
	if month == 2 {
		if isLeapYear {
			return 29
		} else {
			return 28
		}
	} else {
		if month <= 7 && month % 2 ==1 || month >= 8 && month % 2 == 0 {
			return 31
		} else {
			return 30
		}
	}
}

func getDayCount() {
	for {
		fmt.Println("请输入年份：")
		var year int
		_, err := fmt.Scanln(&year)
		if err != nil{
			fmt.Println("年有误")
			continue
		}

		fmt.Println("请输入月份：")
		var month uint8
		_, err2 := fmt.Scanln(&month)
		if err2 != nil {
			fmt.Println("月份有误")
			continue
		}
		
		if 0 >= month || month > 12 {
			fmt.Println("输入的月份有误")
			continue
		}
		fmt.Printf("该月份一共有%v天\n", getMonthDayCount(int(month), judgeIsLeapYear(int(year))))
	}
}

// func main() {
// 	getDayCount()
// }

// 编写一个函数：随机猜数游戏，随机生成一个1-100的数，有十次猜数机会
func guessGame() {
	rand.Seed(time.Now().UnixNano())
	ans := rand.Intn(20)
	times := 0
	fmt.Println("请在[0,100)间猜数：")
	for i := 1; i <= 10; i++ {
		var guess uint8
		fmt.Scanln(&guess)
		if guess == uint8(ans) {
			times = i
			break
		} else {
			fmt.Println("错误请重试")
		}
	}
	if times == 0 {
		fmt.Printf("抱歉您的机会用完了，正确答案是%v", ans)
	} else {
		switch times {
		case 1:
			fmt.Printf("你猜了%v次猜中，你真是个天才", times)
		case 2, 3:
			fmt.Printf("你猜了%v次猜中，你很聪明，赶上我了", times)
		case 4,5,6,7,8,9:
			fmt.Printf("你猜了%v次猜中，一般般", times)
		}
	}
}

func main() {
	guessGame()
}