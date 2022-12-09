package main

import "fmt"

func main() {

	userch := make(chan string, 2)

	// go func() {
	userch <- "Abhi"
	user := <-userch
	fmt.Println(user)
	userch <- "Alice"
	userch <- "foofbasd"
	// }()

	user = <-userch
	fmt.Println(user)
}
