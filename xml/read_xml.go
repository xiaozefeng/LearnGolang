package main

import (
	"encoding/xml"
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	file ,err := os.Open("data.xml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	v := &RecurServers{}
	err = xml.Unmarshal(data, v)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(v)

}
type RecurServers struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
	Description string `xml:",innerxml"`
}

type server struct {
	XMLName xml.Name `xml:"server"`
	ServerName string`xml:"serverName"`
	ServerIp string `xml:"serverIp"`
}

