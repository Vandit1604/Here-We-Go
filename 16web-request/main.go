package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// fmt.Println("We will learn about web side of Golang")

	resp, err := http.Get("http://example.com/")

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	databytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))

	// It's your responsibility to close the connection it won't happen automatically
	defer fmt.Println("Connection closed")
	defer resp.Body.Close()

}
