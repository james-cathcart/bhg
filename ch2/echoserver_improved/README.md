# Echo Server
You can test this by running the application:
```
go run main.go
```
Then connect via Telnet and send any text
```
telnet localhost 20080
```
## Improvement
The improvement is the introduction of the bufio package which automatically handles reading and buffering the data.

**Previous Approach**
```go
// echo is a handler function that simply echoes received data
func echo(conn net.Conn) {
	defer conn.Close()

	// create a buffer to store received data
	b := make([]byte, 512)
	for {

		// receive data via conn.Read into a manual buffer
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

		// send manually buffered data via conn.Write
		log.Println("writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("unable to write data")
		}
	}
}
```
**Improved Approach**
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