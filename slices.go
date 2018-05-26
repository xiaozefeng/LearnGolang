package main

import "fmt"

func main() {

	// 初始化的方式
	// 1.基于数组
	arr:= [10]int{1,2,3,4,5,6}
	sl1:= arr[:]
	fmt.Println("sl1:",sl1)
	sl1 = append(sl1, 7)
	fmt.Println("sl1:",sl1)


	sl2 := make([]int , 5, 10)
	fmt.Println("len(sl2)", len(sl2))
	fmt.Println("cap(sl2)", cap(sl2))

	for _,v := range sl1{
		fmt.Printf("%d,", v)
	}

}
