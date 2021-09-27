package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	health := func(writer http.ResponseWriter, request *http.Request) {
		log.Fatalln(io.WriteString(writer, "OK"))
	}
	mux.HandleFunc("/health", health)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.SetKeepAlivesEnabled(true)
	log.Fatalln(server.ListenAndServe())
}
