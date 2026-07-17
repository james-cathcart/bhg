package main

import (
	"io"
	"log"
	"net"
)

// echo is a handler function that simply echoes received data
func echo(conn net.Conn) {
	defer conn.Close()

	// create a buffer to store received data
	b := make([]byte, 512)
	for {

		// receive data via conn.Read into a buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("client disconnected")
			break
		}
		if err != nil {
			log.Println("unexpected error")
			break
		}
		log.Printf("received %d bytes: %s\n", size, string(b))

		// send data via conn.Write
		log.Println("writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("unable to write data")
		}
	}
}

func main() {

	// bind to tcp port 20080 on all interfaces

	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("unable to bind to port")
	}
	log.Println("listening on 0.0.0.0:20080")
	for {
		// wait for connection. create net.conn on connection established.
		conn, err := listener.Accept()
		log.Println("received connection")
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		// handle the connection using goroutine for concurrency
		go echo(conn)
	}
}
