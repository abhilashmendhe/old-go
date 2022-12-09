package main

import (
	"fmt"
	"log"
)

func socialize() error {
	defer fmt.Println("GoodBye")
	fmt.Println("Hi")
	fmt.Println("How are you")
	return fmt.Errorf("I don't want to talk....")
	return nil
}

func main() {

	err := socialize()
	if err != nil {
		log.Fatal(err)
	}
}
