package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func worker(timeout time.Duration, verbose bool, ports, results chan int) {
	for p := range ports {
		if verbose {
			fmt.Printf("%sscanme.nmap.org%s:%s%d%s\n", ColorCyan, ColorReset, ColorYellow, p, ColorReset)
		}
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {

	var timeoutFlag = flag.Int("t", 5000, "timeout value ms")
	var verboseFlag = flag.Bool("v", false, "verbose output boolean")

	flag.Parse()

	timeout := time.Duration(*timeoutFlag) * time.Millisecond
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	fmt.Println("# Starting scan...")
	for i := 0; i < cap(ports); i++ {
		go worker(timeout, *verboseFlag, ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for range 1024 {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)

	fmt.Println("## --- Scan completed --- ##")
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d %sopen%s\n", port, ColorRed, ColorReset)
	}
}
