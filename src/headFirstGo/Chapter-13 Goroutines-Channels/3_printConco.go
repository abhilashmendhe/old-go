package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}
func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func main() {

	fmt.Println("Start....")
	go a()
	go b()
	time.Sleep(time.Millisecond)
	fmt.Println("\nEnded....")
}
