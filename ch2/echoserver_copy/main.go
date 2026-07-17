package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	// copy data from io.Reader to io.Writer via io.Copy
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("unable to read/write data")
	}
}

func main() {

	// bind to tcp port 20080
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("unable to bind to port")
	}
	log.Println("listening on 0.0.0.0:20080")
	for {
		// wait for connection. create net.Conn on connection established
		conn, err := listener.Accept()
		log.Println("received connection")
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		// handle the connection using goroutine for concurrency
		go echo(conn)
	}
}
