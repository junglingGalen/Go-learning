package main

import (
	"fmt"
)

/* 
接口的快速实现
 */


//定义一个接口（我理解就是定义了一个协议，约束了接入对象所需具有的方法）
type Usb interface {
	Start()
	Close()
}

type Computer struct {

}

// 入参为接口类型，则代表开放接口，可接入满足接口约束的对象
func (c *Computer) Insert(u Usb) {
	u.Start()
	u.Close()
}

type Camera struct {
	// 定义个对象，实现了接口中要求的方法，则可以插入对应的接口
}

func (c Camera) Start() {
	fmt.Println("相机启动了")
}

func (c Camera) Close() {
	fmt.Println("老板，相机关闭了")
}

type Typer struct {

}

func (t Typer) Start() {
	fmt.Println("打印机开始工作了")
}

func (t Typer) Close() {
	fmt.Println("打印机停止工作")
}

/* 接口的注意事项和细节 
	1. 接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的实例
	2. 接口中所有方法都没有方法体，即都没有实现的方法
	3. 在Golang中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了接口
	4. 一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋值给接口类型
	5. 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
	6. 一个自定义类型可以实现多个接口
	7. Golang接口中不能有任何变量
	8. 一个接口(比如A接口)可以继承多个别的接口(比如B,C接口)，这时如果要实现A接口，也必须将B,C接口的方法也全部实现
	9. interface类型默认是一个指针，如果没有对interface初始化就使用，那么会出现nil
	10. 空接口interface{}没有任何方法，所以所有类型都实现了空接口
*/

func main() {
	var c Computer
	var camera Camera
	var typer Typer
	c.Insert(camera)
	c.Insert(typer)
}