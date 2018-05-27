package main

import (
	"time"
	"fmt"
)

// 计算时间差

func main() {

	k := time.Now()
	d, _ := time.ParseDuration("-24h")
	// 一天之前
	fmt.Println(k.Add(d))

	// 一周之前
	fmt.Println(k.Add(d * 7))

	// 一个月之前
	fmt.Println(k.Add(d * 30))


	d1, _ := time.ParseDuration("24h")
	fmt.Println("一天之后",k.Add(d1))
}
