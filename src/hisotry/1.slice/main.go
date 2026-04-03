package main

import (
	"fmt"
)

/*
为什么需要切片：保存数量不固定的数据
切片的基本介绍： 
①切片是数组的一个引用，因此切片是引用数据类型，在进行传递时，
遵守应用的传递机制
② 切片的使用和数组类似，遍历切片、访问切片的元素和求切片的长度都一样
③ 切片的长度是可以变化的，因此切片是一个可以动态变化的数组
④ 切片定义的基本语法：
	var 变量名 []类型
*/

func slice() {
	// 基本使用
	var intArr [5]int = [...]int{1,22,33,44,55}
	// 声明一个切片
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是=", slice)  // 22, 23
	fmt.Println("slice 的元素个数是=", len(slice))  // 2
	fmt.Println("slice 的容量是=", cap(slice))  // 4  切片的容量是随着切片长度的增长自增长的

	// 切片是引用，地址看对应元素和数组相同
	fmt.Println("slice的0号元素的地址是=", &slice[0])  
	fmt.Println("intArr的1号元素地址是=", &intArr[1])  
	// 切片其实就是一个数据结构，它的地址里只记录了自己第一个元素的地址、自身长度、容量


}

// func main() {
// 	slice()
// }

/*切片的定义方式*/
func sliceCreate() {
	// 1.切数组获得切片
	// 2.make([]type, len, [cap])  cap > len
	var slice []float64 = make([]float64, 5, 6)
	slice[2] = 3
	fmt.Println(slice)

	/* 
	方式1和2的区别： 
		①方式1的切片是直接引用了一个已经存在了的数组的部分数据，程序员可以通过数组和切片共同访问这一片数据
		②方式2是通过make来创建切片，make也会创建一个数组，是切片的底层维护，程序员看不见
	*/

	// 3.定义一个切片，直接就指定具体数组，使用原理类似make方式
	var slice2 []string = []string{"Tom", "jsck", "mary"}
	fmt.Println("strSlice=", slice2, len(slice2), cap(slice2))
}

// func main()  {
// 	sliceCreate()
// }

/*切片的遍历*/
func sliceFor() {
	var slice []int64 = make([]int64, 5, 5)
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	for i, v := range slice {
		fmt.Println(i, v)
	}
}

// func main() {
// 	sliceFor()
// }

/*
切片的注意事项和使用细节
	1. var slice = arr[startindex:endindex]
	说明：从arr数组下表为startindex，取到下表为endindex的元素（不含arr[endindex]）
	2. 初始化切片时不能越界。范围在[0-len(arr)]之间，但是可以动态增长
	3.	① var slice = arr[0:end] 可以简写 var slice = arr[:end]
		② var slice = arr[start:len(arr)] 可以简写成 var slice = arr[start:]
		③ var slice = arr[0:len(arr)] 可以简写成 var slice = arr[:]
	4. cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
	5. 切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者make一个空间供切片来使用
	6. 切片可以继续切片
*/


/*切片的append函数，对切片进行动态追加*/
func slice3() {
	var slice3 []int = [] int{100, 200, 300}
	// 通过append直接给slice3追加具体元素
	slice3 = append(slice3, 400, 500)  // 和python不同，这里append有返回值，不修改原切片

	// 通过append将切片slice3添加给slice3
	slice3 = append(slice3, slice3...)  // 注意后面...不能漏，切只能是切片，不能是数组
	
	fmt.Println(slice3)

	var arr[3]int = [3]int{1, 2, 3}
	s := arr[:]
	// s的第一个元素和arr第一个元素相同，修改s的元素，对应arr元素也会修改
	s[2] = 1100
	fmt.Println("扩容前s和arr第一个元素地址：",&s[0], &arr[0])
	fmt.Println("修改s后s和arr：", s, arr)
	// 对s进行扩容后：s的第一个元素和arr第一个元素不再相同，修改s的元素，对应arr元素也不会修改，两者不再是同一段数据
	s = append(s, 4)
	fmt.Println("扩容后s和arr第一个元素地址：", &s[0], &arr[0])
	s[0] = 2200
	fmt.Println("修改s后s和arr：", s, arr)
	// 对s进行扩容，比较扩容前后地址是否发生变化
	fmt.Println("s扩容前地址：", &s[0])
	s2 := append(s, 5050)
	fmt.Println("s扩容后地址：", &s2[0])
	fmt.Println("s和扩容后的s2：", s, s2)
	s2[1] = 2360
	fmt.Println("修改s2后比较s和s2：", s, s2)
	/*
	事实证明，扩容后切切片会创建一个新的数组来追加具体元素,
		① 如果旧切片属于某一个已定义数组的一部分，则新切片地址不再和旧切片相同，修改新切片，旧切片和数组不再变化
		② 如果旧切片不属于某一个已定义数组的一部分，则新切片地址和旧切片相同，修改新切片，旧切片对应值也会同样改变
	*/
}

// func main() {
// 	slice3()
// }

func copslice() {
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 []int = []int{2:100, 10:0}
	fmt.Println("拷贝前：", slice4, slice5)
	copy(slice5, slice4)  // slice5 = [1, 2, 3, 4, 5, 0, 0, 0, 0, 0]
	fmt.Println("拷贝后：", slice4, slice5)
	var slice6 []int = make([]int, 3)
	i := copy(slice6, slice4)  // slice6 = [1, 2, 3]
	fmt.Println("slice6:", slice6, i)
	/*
	① copy的数据类型是切片
	② slice4 和 slice5 的数据空间是独立的，互不影响
	③copy(slice5, slice4): 把slice4对应下表的元素赋值给slice5对应下标。若超出slice5现在长度，则超出部分不赋值
	*/

}

func main() {
	copslice()
}