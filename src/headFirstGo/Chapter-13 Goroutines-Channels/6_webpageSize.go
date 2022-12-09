package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, channel chan Page) {

	// fmt.Print("Getting URL: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf(" size if %0.2f KB\n", float64(len(body)/1000))
	channel <- Page{URL: url, Size: len(body)}
}

func main() {

	webPages := []string{"http://example.com/", "https://golang.org/", "https://golang.org/doc", "https://www.youtube.com"}

	pages := make(chan Page)

	for _, url := range webPages {
		go responseSize(url, pages)
	}

	for i := 0; i < len(webPages); i++ {
		page := <-pages
		fmt.Println(page.URL, page.Size)
	}
}
