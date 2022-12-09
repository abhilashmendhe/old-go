package main

import (
	"fmt"
)

func greeting(myChannel chan string) {
	myChannel <- "hi how are you"
}
func main() {

	myChannle := make(chan string)
	go greeting(myChannle)
	receivedVal := <-myChannle
	fmt.Println(receivedVal)
}
