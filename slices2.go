package main

import "fmt"

func main() {
	s1:= []int {1,2,3,4,5}
	s2:= []int {7,8,9}

	// 将s2拷贝到s1
	copy(s1,s2)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)


	s3:= []int {1,2,3,4,5}
	s4:= []int {7,8,9}

	// 将s2拷贝到s1
	copy(s4,s3)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

}
