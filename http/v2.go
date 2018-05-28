package main

import (
	"net/http"
	"log"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &MyHandler{})
	mux.HandleFunc("/bye", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Bye bye, this is version 2!"))
	})
	log.Println("Start Server At locahost:9000 ...")
	log.Fatal(http.ListenAndServe(":9000", mux))
}

type MyHandler struct {
}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, this is version 2!"))
}
