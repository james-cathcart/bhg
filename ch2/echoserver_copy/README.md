# Echo Server
You can test this by running the application:
```
go run main.go
```
Then connect via Telnet and send any text
```
telnet localhost 20080
```
## IO Copy Approach
Simplify the approach even further with Go's io.Copy

**Previous Approach**
```go
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
```

**IO Copy**

```go
func echo(conn net.Conn) {
	defer conn.Close()

	// copy data from io.Reader to io.Writer via io.Copy
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("unable to read/write data")
	}
}
```