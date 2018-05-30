package main

import (
	"os/exec"
	"os"
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	//sl := make([]string, 0, 10)
	Parent := "/Users/xiaozefeng/work/20170814/Parent"
	Gateway := "/Users/xiaozefeng/work/20170814/Gateway-bt"
	Underground := "/Users/xiaozefeng/work/20170814/Underground-bt"
	Facilities := "/Users/xiaozefeng/work/20170814/BTFacilities"
	Manager := "/Users/xiaozefeng/work/20170814/BTCreditManager"
	Market := "/Users/xiaozefeng/work/20170814/BTCreditMarket"

	//sl = append(sl, Parent)
	//sl = append(sl, Gateway)
	//sl = append(sl, Underground)
	//sl = append(sl, Facilities)
	//sl = append(sl, Manager)
	//sl = append(sl, Market)

	arr := []string{
		Parent,
		Gateway,
		Underground,
		Facilities,
		Manager,
		Market,
	}

	//sl1 := sl[:3]
	for _, v := range arr {
		str := fmt.Sprintf("cd %s && mvn clean install -Dmaven.test.skip=true", v)
		fmt.Println(str)
		cmd := exec.Command("/bin/sh", "-c", str)
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
	}

	//sl2 := sl[3:]
	////ch := make(chan int, 3)
	//for _, v := range sl2 {
	//	str := fmt.Sprintf("cd %s && mvn clean install -Dmaven.test.skip=true", v)
	//	//go func() {
	//	cmd := exec.Command("/bin/sh", "-c", str)
	//	cmd.Stdout = os.Stdout
	//	err := cmd.Run()
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		os.Exit(0)
	//	}
	//	//ch <- 1
	//	//}()
	//}
	////<-ch
	result := time.Since(start)
	fmt.Println("花费了多少秒", result)

}
