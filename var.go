package main

import "fmt"

var a int
var arr [3]int
var sl []int
var st struct{}
var p *int
var f func(a int) int
var (
	v1 int
	v2 string
)

func main() {
	fmt.Printf("int:%d\n", a)
	fmt.Printf("arr:%v\n", arr)
	fmt.Printf("slice:%v\n", sl)
	fmt.Printf("struct:%v\n", st)
	fmt.Printf("point:%v\n", p)
	fmt.Printf("function:%s\n", f)

	// 变量赋值方式
	var v3 = 123
	var v4, v5 = 1, 2
	v6 := 3
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)
	fmt.Println(v6)

	// 交互变量的值
	v4, v5 = v5, v4
	fmt.Println(v4)
	fmt.Println(v5)

}
