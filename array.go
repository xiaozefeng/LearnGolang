package main

import "fmt"

// 在go语言中数组是值类型
func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("arr:", arr)
	// 遍历
	fmt.Println("遍历1:")
	for i := 0; i<len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("--------------")

	for  index, value := range arr {
		fmt.Printf("index:%d: value: %d\n", index, value)
	}


	modifyArray(arr)
	// 发现数组并没有改变
	fmt.Println(arr)


}

func modifyArray(arr [10]int)  {
	arr[0] = 10
	fmt.Println("arr:",arr)
}
