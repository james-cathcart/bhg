package main

import (
	"fmt"
	"log"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {
	// instantiate reader and writer
	var (
		reader FooReader
		writer FooWriter
	)

	// create buffer to hold input/output
	input := make([]byte, 4096)

	// use reader to read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	// use writer to write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("unable to write data")
	}
	fmt.Printf("wrote %d bytes\n", s)
}
