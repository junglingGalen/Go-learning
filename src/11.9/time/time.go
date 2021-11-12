package main

import (
	"time"
	"fmt"
)

func main() {
	// 看看日期和时间相关函数和方法的使用
	now := time.Now()
	fmt.Printf("now=%v type=%T", now, now)

	// 通过now获取到年月日 时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", int(now.Month())) // 返回的是time.Month类型，用int装换为数字
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 格式化日期时间
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(),
	 now.Hour(), now.Minute(), now.Second())
	
	 dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(),
	 now.Hour(), now.Minute(), now.Second())
	 fmt.Println("dateStr=", dateStr)
	// 格式化第二种方法
	fmt.Printf("%T\n", now.Format("2006/01/02 15:04:05"))  // format的字符串是固定的，不要问我为什么
	// 备注：各个数字可以随机组合

	// 时间常量
	// for i:=1; i<10; i++ {
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second * 2)  // 这里传入必须是整数，如果是小数需要转化成更低级别的时间单位乘以整数
	// }

	// time的Unix 和UnixNano
	fmt.Printf("unix时间戳=%v, unixnano时间戳=%v", now.Unix(), now.UnixNano())

}