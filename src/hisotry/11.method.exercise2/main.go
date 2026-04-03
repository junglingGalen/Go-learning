package main

import (
	"fmt"
)

/*编写方法：判断一个数是奇数还是偶数*/

type integer int

func (i integer) isOdd() string{
	if i % 2 == 0 {
		return "Odd"
	} else {
		return "notOdd"
	}
}

/*根据行、列、字符打印对应行数和列数的字符，比如： 行：3，列：2.字符*,则打印对应效果 */

func (i integer) print(m int, n int, k string) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%v", "*")
		}
		fmt.Println(" ")
	}
}

/*编写一个方法，返回3*3二维数组的转质*/
type Arr [3][3]int

func (a Arr) T() [3][3]int {
	var temArr [3][3]int
	for ind, _arr := range a {
		for col, val := range _arr {
			temArr[col][ind] = val
		}
	}
	return temArr
}

func main(){
	var arr1 = [3][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	arr2 := Arr(arr1)
	fmt.Println(arr2.T())
}

