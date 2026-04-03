package main

import (
	"fmt"
	"sort"
)

/*
① map的基本语法
var map 变量名 map[keytype]valuetype

② key可以是什么类型：
bool,数字，string，指针，channel,还可以是只包含前面几个类型的 接口、结构体、数组

注意：slice、map 还有function不可以，因为这几个没法用 == 来判断

③ value可以是什么类型：
和key基本一样

④ 声明map是不会分配内存的，初始化需要make，分配内存后才能赋值使用
*/

func mapFunc1() {
	// map的声明
	var a map[string]string
	a = make(map[string]string, 1)  //给map分配空间, 默认长度为0的空间
	a["nol"] = "宋江"
	a["no2"] = "吴用"
	a["no2"] = "武松"  // map的key不能重复,相同key会覆盖
	a["no3"], a["no4"]= "李逵", "武大郎"
	fmt.Println(a)

	/* map声明的3种方式 */

	// ① 先声明，后分配空间如上

	// ② 声明时同时分配空间
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "东京"

	// ③声明时空时分配空间，并赋值元素
	heroes := map[string]string{
		"hero1" : "孙悟空",
		"hero2": "卢俊义",  // 这个地方,不能少（和python不一样！）
	}

	fmt.Println(cities, heroes)
}

// func main() {
// 	mapFunc1()
// }

func mapFunc2() {
	/*建立个map里面存放学生信息，每个学生有名字和性别两个字段*/
	var studentInfo map[string]map[string]string = make(map[string]map[string]string, 0)

	studentInfo["lili"] = make(map[string]string, 0)
	studentInfo["lili"]["name"] = "lili"
	studentInfo["lili"]["sex"] = "girl"

	studentInfo["Tom"] = map[string]string{"name": "Tom", "sex": "boy"}

	fmt.Println(studentInfo)
}

// func main() {
// 	mapFunc2()
// }

func mapFunc3() {
	/*map的增删改查*/
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "东京"

	// ① 增加/修改
	cities["no2"] = "上海~" // 存在则修改，不存在则增加

	// ② 查找
	val, ok := cities["no5"]  // 存在返回对应value和bool； 不存在返回空字符串和false
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("不存在")
	}
	

	// ③删除
	delete(cities, "no3")
	delete(cities, "no3")  //重复删除也不报错
	// 若要清空map的所有key： ①遍历所有的key逐一删除； ②直接make一个新的空间，原空间会被垃圾回收
	cities = make(map[string]string)
	fmt.Println(cities)

}

// func main() {
// 	mapFunc3()
// }

func mapFunc4() {
	/*map的遍历*/
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "东京"

	//使用for-range遍历
	for k, v := range cities {
		fmt.Println("k=", k, "v=", v)
	}
}

// func main() {
// 	mapFunc4()
// }

func mapFunc5() {
	/* slice联合map使用 */
	slice := make([]map[string]string, 2)
	
	if slice[0] == nil {
		slice[0] = map[string]string{
			"name": "哥斯拉",
			"age": "24",
		}
	}

	if slice[1] == nil {
		slice[1] = map[string]string{
			"name": "白骨精",
			"age": "16",
		}
	}

	// if slice[2] == nil {
	// 	slice[2] = map[string]string{
	// 		"name": "玉兔精",
	// 		"age": "14",
	// 	}
	// }       ----------> 会报错，超出了slice的长度范围，要改成用append方法才可以动态增长长度

	var monster map[string]string
	monster = make(map[string]string, 2)
	monster["name"] = "玉兔精"
	monster["age"] = "14"
	slice = append(slice, monster)

	fmt.Println(slice)
}

// func main() {
// 	mapFunc5()
// }

func mapFunc6() {
	/*map的排序，map是无序的，并且每次遍历得到的顺序也不一定一致的*/
	map1 := make(map[int]int, 2)
	map1[10] = 1000
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	/* 如果要按照mao的key的顺序进行排序后输出
		① 将map的key放入到切片中
		② 对切片排序
		③ 遍历切片，然后按照key来输出map
	*/
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	// 排序
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println(map1[v])  // ？这里有点奇怪map1[v]返回的是value，bool,为什么打印出来只有value
	}
}

// func main() {
// 	mapFunc6()
// }

func mapFunc7() {
	/* map的使用细节和陷阱 
	① map是引用类型，遵守引用类型的传递机制
	② map的容量达到后，在想map增加元素，会自动扩容，并不会发生panic。而切片不可以，需要用append
	③ map的value也经常使用struct类型，更适合管理复杂的数据
	*/

	// 定义一个结构体
	type Stu struct {
		Name string
		Age int
		Address string
	}

	students := make(map[string]Stu, 10)
	// 创建两个学生
	stu1 := Stu{"Tom", 18, "北京"}
	stu2 := Stu{"Json", 13, "上海"}

	students["no1"] = stu1
	students["no2"] = stu2

	fmt.Println(students)

	// 遍历各个学生信息
	for k, v := range students {
		fmt.Printf("学生的编号是%v\n", k)
		fmt.Printf("学生的名字是%v\n", v.Name)
		fmt.Printf("学生的年龄是%v\n", v.Age)
		fmt.Printf("学生的地址是%v\n", v.Address)
	}

}

// func main() {
// 	mapFunc7()
// }

func mapFunc8(users map[string]map[string]string, name string) {
	/* 课堂练习
	① 使用 map[string]map[string]string的map类型
	② key：标识用户名，是唯一的，不可以重复
	③ 如果某个用户名存在，就将其密码修改“888888”，如果不存在就增加这个用户信息（包括昵称nickname和密码pwd）
	④ 编写一个函数modifyUser(users map[string]map[string], name string)完成上述功能
	*/
	_, ok := users[name]
	if ok {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = map[string]string{"nickname": name, "pwd": "111111"}
	}
}

func main() {
	users := make(map[string]map[string]string, 2)

	user1 := make(map[string]string, 2)
	user1["nickname"] = "lili"
	user1["pwd"] = "123456"
	users[user1["nickname"]] = user1

	users["Tom"] = map[string]string{"nickname": "Tom", "pwd": "489657"}
	fmt.Println(users)

	mapFunc8(users, "Jack")
	fmt.Println(users)

	mapFunc8(users, "Tom")
	fmt.Println(users)
}