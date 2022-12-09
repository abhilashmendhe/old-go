package main

import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	messages := []byte("Hello, World!")
	_, err := writer.Write(messages)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:8001", nil)
	log.Fatal(err)
}
