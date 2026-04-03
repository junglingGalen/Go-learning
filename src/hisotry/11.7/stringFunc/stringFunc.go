package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//1.统计字符串的长度（按字节数返回）len(str)
	str := "start北"
	fmt.Printf("%T-%v\n", len(str), len(str))

	// 字符串遍历，同时处理有中文的问题 r := []rune(str)
	str2 := "hello北京"
	r := []rune(str2)
	fmt.Printf("%T\n", r)
	for i:=0; i < len(r); i++ {
		fmt.Printf("字符=%c,type=%T\n", r[i], r[i])
	}

	// 字符串转整数：n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("p")
	if err != nil {
		fmt.Println("字符串转换int成功：", n)
	}
	fmt.Println(n, err)

	// 整数转字符串 str = strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v, type=%T \n", str, str)
	fmt.Println(string(9798))  // 该方法是把utf8编码转换成字符

	// 字符串转[]byte: var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	// []byte 转字符串：str = string([]byte{97,98,99})
	str = string([]byte{97,98,99})
	fmt.Printf("str=%v\n", str)

	// 10进制转2,8,16进制字符串： str = strconv.FormatInt(123, 2), 返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是%v - type:%T\n", str, str)

	//查找子串是否在指定的字符串中：strings.Contains("seafood", "foo") //true
	b := strings.Contains("seafood", "foo")
	fmt.Println(b)

	// 统计一个字符串有几个指定的字符串：strings.Count("chiness", "s") //2
	num := strings.Count("chiness", "s")
	fmt.Println(num)

	// 不区分大小写的字符串比较：fmt.Println(strings.EqualFold("abc", "AbC"))
	fmt.Println(strings.EqualFold("abc", "AbC"))

	//返回子串在字符第一次出现的index，没有返回-1：
	index := strings.Index("NLT_北abc_abc", "abc")
	fmt.Println(index)

	//返回子串在字符最后一次出现的index，没有返回-1：
	index = strings.LastIndex("NLT_北abc_abc", "abc")
	fmt.Println(index)

	// 将指定子串替换成另外一个子串： strings.Repalce
	fmt.Println(strings.Replace("123456!!!!", "!", "?", 2)) // -1代表全部换

	// 按照指定的某个字符为分割标识，将一个字符串拆分出字符串数组
	fmt.Println(strings.Split("hello,world,go", ","))

	// 将字符串的字母进行大小写的转换
	fmt.Println(strings.ToUpper("HelloWorld"))
	fmt.Println(strings.ToLower("HelloWorld"))

	// 将字符串左右两边的空格去掉
	fmt.Println(strings.TrimSpace("  I have space "))

	// 将字符串两边指定的字符串去掉
	fmt.Println(strings.Trim("!  trim  ! ", " !"))
	// 去掉左边/右边的字符  TrimLeft/TrimRight

	// 判断字符串是否已指定的字符串开头/结尾
	fmt.Println(strings.HasPrefix("http://www.baidu.com", "http"))
	fmt.Println(strings.HasSuffix("http://www.baidu.com", ".com"))

}