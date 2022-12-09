package main

import (
	"fmt"
)

func calmDown() {
	fmt.Println("Calming down")
	recover()
}
func freakOut() {

	defer calmDown()
	/* defer recover()--> don't use this,,, instead call above func so it will sucessfully complete all
	func calls*/
	panic("oh no")

}

func main() {

	freakOut()
	fmt.Println("Exitinig Normally")
}
