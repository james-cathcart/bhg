package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 1024; i++ {
		wg.Add(1) // increment goroutine counter
		go func(j int) {
			defer wg.Done() // decrement goroutine counter
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
}
