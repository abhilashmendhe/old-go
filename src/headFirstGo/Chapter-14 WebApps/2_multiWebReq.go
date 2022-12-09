package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello English")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut...")
}

func hindiHandler(writer http.ResponseWriter, reqest *http.Request) {
	write(writer, "Namaste Bhaisaab")
}

func main() {

	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
