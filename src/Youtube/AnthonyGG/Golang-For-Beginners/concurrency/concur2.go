package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen() {

running: // you can name your for loop
	for {
		select {
		// block here until someone is sending message to this channel
		case msg := <-s.msgch:
			fmt.Printf("received message from %s; payload: %s\n", msg.From, msg.Payload)
		case <-s.quitch:
			fmt.Println("the server is doing a gracefull shutdown")
			// Logic for the gracefull down
			break running
		default:

		}
	}
	fmt.Println("Server has shutdown!!")
}
func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "Abhilash",
		Payload: payload,
	}
	msgch <- msg
}

func graceFullQuitServer(quitch chan struct{}) {
	close(quitch)
}
func main() {

	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}
	go s.StartAndListen()

	// for i := 0; i < 10; i++ {
	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgch, "Hello Programmer")
	}()
	// }

	go func() {
		time.Sleep(4 * time.Second)
		graceFullQuitServer(s.quitch)
	}()
	select {}
}
