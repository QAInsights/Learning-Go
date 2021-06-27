package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HTTP Tester")

	url := "https://example.com"
	// rps := 5
	// threads := 2
	// duration := 10 // seconds

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Err")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
