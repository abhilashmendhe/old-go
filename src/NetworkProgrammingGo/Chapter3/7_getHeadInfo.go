package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckError(err)

	r, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	CheckError(err)

	fmt.Println(r)
	result, err := ioutil.ReadAll(conn)
	CheckError(err)

	fmt.Println(string(result))
	os.Exit(0)
}
