package main

import (
	"fmt"
	"log"
	"net"
)

type Message struct {
	from    string
	payload []byte
}

type ServeR struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan Message
}

func NewServer(listenAddr string) *ServeR {
	return &ServeR{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Message, 10),
	}
}

func (s *ServeR) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()
	<-s.quitch
	close(s.msgch)
	return nil
}

func (s *ServeR) acceptLoop() {

	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		fmt.Println("New connection to the server:", conn.RemoteAddr())
		go s.reaadLoop(conn)
	}
}

func (s *ServeR) reaadLoop(conn net.Conn) {

	defer conn.Close()
	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
			continue
		}
		// msg := buf[:n]
		// fmt.Println(string(msg))
		s.msgch <- Message{
			from:    conn.RemoteAddr().String(),
			payload: buf[:n],
		}
		conn.Write([]byte("thank you for your message!"))
	}
}

func main() {

	server := NewServer(":3000")
	go func() {
		for msg := range server.msgch {
			fmt.Printf("received message from connection (%s):%s\n", msg.from, string(msg.payload))
		}
	}()
	log.Fatal(server.Start())
}
