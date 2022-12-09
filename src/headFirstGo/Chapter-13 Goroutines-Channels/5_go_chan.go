package main

import "fmt"

func abc(ch chan string) {
	ch <- "a"
	ch <- "b"
	ch <- "c"
}

func def(ch chan string) {
	ch <- "d"
	ch <- "e"
	ch <- "f"
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go abc(ch1)
	go def(ch2)

	fmt.Print(<-ch1)
	fmt.Print(<-ch1)
	fmt.Print(<-ch1)
	fmt.Print(<-ch2)
	fmt.Print(<-ch2)
	fmt.Print(<-ch2)
	fmt.Println()
}
