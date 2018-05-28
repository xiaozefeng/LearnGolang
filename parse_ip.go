package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.ParseIP("127.0.0.1"))
}
