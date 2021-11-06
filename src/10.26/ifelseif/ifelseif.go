package main
import (
	"fmt"
)

func main(){
	var score float64
	fmt.Println("请输入考试成绩")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("你将获得一辆BMW")
	} else if score >= 80 && score < 100 {
		fmt.Println("奖励你一个iPhone7")
	} else if score >= 60 && score < 80 {
		fmt.Println("奖励你一个iPad")
	} else {
		fmt.Println("你考的太差了！")
	}
}