package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Make the request
	response, err := http.Get("http://www.example.com")
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
