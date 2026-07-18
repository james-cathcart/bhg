package main

import (
	"fmt"
	"log"
	"net/http"
)

type router struct {
}

func (rtr *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/a":
		fmt.Fprint(w, "Executing /a")
	case "/b":
		fmt.Fprint(w, "Executing /b")
	case "/c":
		fmt.Fprint(w, "Executing /c")
	default:
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}
}

func main() {
	var r router
	log.Println("starting server at localhost:8000")
	http.ListenAndServe(":8000", &r)
}
