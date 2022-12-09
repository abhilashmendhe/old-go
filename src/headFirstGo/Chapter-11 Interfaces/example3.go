package main

import "fmt"

type Whistle string
type Horn string
type Robot string

type NoiseMaker interface {
	Makesound()
	Walk()
}

func (r Robot) Makesound() {
	fmt.Println("I am Robot 2.0")
}
func (w Whistle) Makesound() {
	fmt.Println(w, "Tweet!")
}
func (h Horn) Makesound() {
	fmt.Println(h, "Honk!")
}

func (r Robot) Walk() {
	fmt.Println("I can walk")
}

func play(n NoiseMaker) {
	n.Makesound()
	n.Walk()
}
func main() {
	// var toy NoiseMaker
	// toy = Whistle("Toyco")
	// toy.Makesound()
	// toy = Horn("Yowe")
	// toy.Makesound()

	// play(Horn("YX"))
	// play(Whistle("ABC"))
	// r := Robot("chitti")
	// r.Makesound()
	// play(r)
	// r.Walk()

	var toy NoiseMaker
	toy = Robot("kai mante")
	play(toy)
}
