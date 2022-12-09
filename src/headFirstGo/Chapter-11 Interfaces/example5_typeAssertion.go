package main

import "fmt"

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("bhaag Bhaag!")
}

type NoiseMaker interface {
	MakeSound()
}

func main() {

	var noisemake NoiseMaker = Robot("chitti")
	noisemake.MakeSound()
	var robot Robot = noisemake.(Robot)
	robot.Walk()
}
