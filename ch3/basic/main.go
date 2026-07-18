package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	/*
	 * Get request
	 */
	resp, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()

	/*
	 * Head request
	 */
	resp, err = http.Head("https://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)

	/*
	 * Form request
	 */
	form := url.Values{}
	form.Add("foo", "bar")
	resp, err = http.Post(
		"https://www.google.com/robots.txt",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)

	/*
	 * Delete request
	 */
	req, err := http.NewRequest(http.MethodDelete, "https://www.google.com/robots.txt", nil)
	if err != nil {
		log.Panicln(err)
	}
	var client http.Client
	resp, err = client.Do(req)
	resp.Body.Close()
	fmt.Println(resp.Status)

	/*
	 * Head request
	 */
	req, err = http.NewRequest(
		http.MethodPut,
		"https://www.google.com/robots.txt",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Panicln(err)
	}
	resp, err = client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)

}
