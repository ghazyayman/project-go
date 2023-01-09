package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Create the request body
	body := []byte("Hello, World!")
	req, err := http.NewRequest("POST", "http://www.example.com", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the content type header
	req.Header.Set("Content-Type", "text/plain")

	// Make the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	// Read the response
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contents))
}
