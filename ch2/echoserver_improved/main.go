package main

import (
	"bufio"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("unable to read data")
	}
	log.Printf("read %d bytes: %s", len(s), s)

	log.Println("writing data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("unable to write data")
	}
	writer.Flush()
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
