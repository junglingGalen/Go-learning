package main

import (
	"fmt"
	"strings"
)

func main() {
	a := '中'
	fmt.Printf("值：%v 类型：%T\n", a, a)

	str := "asdfasdf"
	for _, v := range str {
		fmt.Printf("%v(%c)", v, v)
	}
	fmt.Println()

	// ========== 字符串替换 ==========

	s1 := "zhongguo"

	// 方式1: strings.ReplaceAll — 简单替换，最常用
	s2 := strings.ReplaceAll(s1, "guo", "wen")
	fmt.Println("ReplaceAll:", s2) // zhongwen

	// 方式2: strings.NewReplacer — 多组替换，一次遍历完成，性能最好
	replacer := strings.NewReplacer("zhong", "中", "guo", "国")
	s3 := replacer.Replace(s1)
	fmt.Println("NewReplacer:", s3) // 中国

	// 方式3: []byte 转换 — 修改单个ASCII字符，零拷贝思路
	bs := []byte(s1)
	bs[0] = 'Z'
	fmt.Println("[]byte修改:", string(bs)) // Zhongguo

	// 方式4: []rune 转换 — 修改中文等多字节字符
	s4 := "你好世界ad"
	rs := []rune(s4)
	rs[2] = '中'
	fmt.Println("[]rune修改:", string(rs)) // 你好中界

	fmt.Println("%v", []rune(s4))
}
