package main

import (
	"net/http"
	"time"
	"fmt"
	"log"
)

func main() {
	server := &http.Server{
		Addr:":9000",
		WriteTimeout:2 * time.Second,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello this is version 3 !"))
	})
	mux.HandleFunc("/bye", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Bye byem this is version 3!"))
	})

	server.Handler = mux
	fmt.Println("Started Server At localhost:9000")
	log.Fatal(server.ListenAndServe())
}
