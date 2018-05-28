package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s:= marshal()
	fmt.Println(s)
	unmarshal(s)

}

func marshal() string{
	var s ServerSlice
	s.Servers = append(s.Servers, Server{ServerName:"server1", ServerIp:"127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName:"server2", ServerIp:"127.0.0.1"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(b)
}

func unmarshal(str string) {
	var s ServerSlice
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

type Server struct {
	ServerName string
	ServerIp   string
}
type ServerSlice struct {
	Servers []Server
}
