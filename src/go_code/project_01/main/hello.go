//在go中每一个文件都需要归属于一个包
package main  
import "fmt"  //引入一个包，包名 fmt, 引入该包后就可以使用fmt包的函数；（这里和python不同，是字符串）
func main() {  //func是一个关键字，表示一个函数，main是函数名，是一个主函数，是程序的入口
	fmt.Println("hello,world!")
}