package main

import (
	"math"
	"fmt"
)

const PI float64 = math.Pi
const (
	Sunday=iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays
)

const (
	v1 =iota
	v2
	v3
)

const(
	var1 = 1<<iota
	var2
	var3
)

func main() {

	fmt.Println("PI:",PI)

	fmt.Println("Sunday:",Sunday)
	fmt.Println("Monday:",Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:",Thursday)
	fmt.Println("Friday:",Friday)
	fmt.Println("Saturday:", Saturday)
	fmt.Println("NumberOfDays:", numberOfDays)

	fmt.Println("------------------")

	fmt.Println("v1:",v1)
	fmt.Println("v2:",v2)
	fmt.Println("v3:",v3)

	fmt.Println("------------------")

	fmt.Println("var1:",var1)
	fmt.Println("var2:",var2)
	fmt.Println("var3:",var3)





}
