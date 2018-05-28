package main

import (
	"net/http"
	"time"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	server := &http.Server{
		Addr:         ":9000",
		WriteTimeout: 2 * time.Second,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello this is version 4 !"))
	})
	mux.HandleFunc("/bye", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Bye byem this is version 4!"))
	})

	go func() {
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close Server ", err)
		}
	}()

	server.Handler = mux
	fmt.Println("Started Server At localhost:9000")
	err := server.ListenAndServe()
	if err == http.ErrServerClosed {
		log.Println("Server Closed under request")
	} else {
		log.Fatal("Server Closed unexpected")
	}
	log.Println("Server Exit")
}
