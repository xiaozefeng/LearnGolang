package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println(r.URL.Path)
		fmt.Println(r.URL.Scheme)
		fmt.Fprintf(writer, "Hello, World")
	})
	err:= http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Fatal("error", err.Error())
	}

}
