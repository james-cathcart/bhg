package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get(`name`))
}

func main() {

	http.HandleFunc(`/hello`, hello)
	log.Println("starting server at localhost:8080")
	http.ListenAndServe(":8080", nil)

}
