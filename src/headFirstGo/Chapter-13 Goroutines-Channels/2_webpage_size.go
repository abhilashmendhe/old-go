package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {

	fmt.Print("Getting URL: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" size if %0.2f KB\n", float64(len(body)/1000))
}

func main() {

	go responseSize("http://example.com/")
	go responseSize("https://golang.org/")
	go responseSize("https://golang.org/doc")
	go responseSize("https://www.youtube.com")

	time.Sleep(time.Second)
}
