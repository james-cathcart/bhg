package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("start")
	l.Inner.ServeHTTP(w, r)
	log.Println("finish")
}

type authenticator struct {
	Inner    http.Handler
	Username string
	Password string
}

func (a *authenticator) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println("checking auth...")

	// get query param values
	username := r.URL.Query().Get(`u`)
	password := r.URL.Query().Get(`p`)

	// do auth comparison
	if username != a.Username || password != a.Password {

		log.Printf("error: unauthenticated")

		log.Printf("\tusername value: %s, input: %s\n", a.Username, username)
		log.Printf("\tpassword value: %s, input: %s\n", a.Password, password)

		http.Error(w, "Unauthenticated", http.StatusUnauthorized)

		return
	}

	// logged in
	log.Printf("%s logged in", username)
	ctx := context.WithValue(r.Context(), "username", username)
	r = r.WithContext(ctx)

	// call inner handler
	a.Inner.ServeHTTP(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello\n")
}

func main() {

	helloHandler := http.HandlerFunc(hello)

	// create middleware chain
	mux := logger{
		Inner: &authenticator{
			Inner:    helloHandler,
			Username: `foo`,
			Password: `bar`,
		},
	}

	log.Println("starting server at localhost:8000")
	http.ListenAndServe(":8000", &mux)
}
