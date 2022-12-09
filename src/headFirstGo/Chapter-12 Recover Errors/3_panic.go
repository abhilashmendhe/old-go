package main

import "fmt"

func one() {
	defer fmt.Println("one...")
	two()
}
func two() {
	defer fmt.Println("two...")
	three()
}
func three() {
	defer fmt.Println("three...")
	panic("wow..")
}
func main() {
	one()
}
