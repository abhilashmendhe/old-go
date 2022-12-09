package main

import (
	"interfaces/example1/gadget"
)

func playList(device gadget.TapePlayer, songs []string) {

	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {

	player := gadget.TapePlayer{Batteries: "YES"}
	mixtape := []string{"Aadat", "Dil Harey", "Bulleya", "Tum hi ho"}
	playList(player, mixtape)

	playe = gadget.TapeRecorder{}
	playList(playe, mixtape)
}
