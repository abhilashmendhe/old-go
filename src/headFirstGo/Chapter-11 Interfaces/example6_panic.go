package main

import "fmt"

type Player interface {
	Play(string)
	Stop()
}

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("tapePlayer: Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("tapeRecorder: Playing", song)
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

func TryOut(player Player) {
	player.Play("Mettalica")
	player.Stop()

	recorder, ok := player.(TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player is not recorder, therefore can't record!!!")
	}
}

func main() {
	TryOut(TapeRecorder{Microphones: 1000})
	TryOut(TapePlayer{Batteries: "Hai"})
}
