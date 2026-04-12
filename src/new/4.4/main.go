package main

import (
	"fmt"
	"sync"
)

func _main() {
	var a int = 10
	switch {
	case a <= 10:
		fmt.Println("<100")
	case a<=100:
		fmt.Println("<1000")
	default:
		fmt.Println("....")
	}
}


func __main(){
	a := []int{1,2,3,4,5}
	fmt.Println(a)
	b := a[3:]
	fmt.Println(b)
	c := append(b, 3)
	fmt.Println(c)
	c[1] = 100
	fmt.Println(b, c)
	fmt.Println(cap(c), len(c))
	fmt.Println(cap(b), len(b))

	var d []int
	fmt.Println(d)
	fmt.Println(cap(d), len(d))

	var e = make([]int, 4, 5)
	fmt.Println(e)
	fmt.Println(cap(e), len(e))

	f := e
	f[1]=3
	fmt.Println(e)

	g := []int{1,2,3}
	g = append(g, g...)
	fmt.Println(g)

	h := []int{1,2,3}
	i := append(h, 4)
	i[1]=1000
	j := append(i, 5)
	j[2]=2000
	fmt.Println(h, i, cap(h), cap(i), j, cap(j))
}

func ___main(){
	var wg sync.WaitGroup
	
	m := make(map[int]int)
	// 启动100个goroutine同时写
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m[i] = i  // ← 多个goroutine同时执行这里 → panic!
		}(i)
	}
	wg.Wait()

}

func ____main(){
	dict := []map[string]int{
		{"lili": 1,
		"lanlan": 2,},
	}
	fmt.Println(dict[0]["lildi"])

	d := make([]int, 2, 5)
	g := []int{1,2,3,4,5}
	f := append(d, g...)
	f[0] = 9
	fmt.Println(d, f)
	l := g
	l = append(l, g...)
	l[0] = 100
	fmt.Println(l, g)

	k := []int{1,2}
	p := append(k, 3)
	fmt.Println(k, p)

}

func main_(){
	var i = 1
	defer func(){fmt.Println(i)}()
	i ++
}


func main(){
	var ch chan int
	ch = make(chan int)
	ch <- 1
	fmt.Println("1")
	ch <- 1
	fmt.Println("2")
	ch <- 1
	fmt.Println("3")
	ch <- 1
	fmt.Println("4")
	<- ch
	fmt.Println(5)
}