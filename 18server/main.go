package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myUrl string = "http://localhost:8000"

// GET - This will request for something to the URL
// POST - This will post/give something to the URL. This sends `application/json` format
// POSTFORM - This sends encoded form format

func main() {
	fmt.Println("Let's create a GET,POST and POSTFORM request")
	// PerformGetRequest(myUrl)
	// PerformPostJsonRequest(myUrl)
	PerformPostFormRequest(myUrl)
}

// Sends a POST request to the server. Our server is configured to send back the sent response.
// POST request sends a JSON
func PerformPostJsonRequest(myUrl string) {
	myUrl = myUrl + "/post"
	// how to create fake json payload
	requestBody := strings.NewReader(
		`
	{
		"name":"vandit",
		"age":12,
		"masti":true
	}
	`)

	// content type should be that when we're sending JSON
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

// Sends a GET request to the server. this will return this in response: Hello from localhost:8000
func PerformGetRequest(myUrl string) {
	// update for get request
	myUrl = myUrl + "/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	// always the caller should close the response
	defer response.Body.Close()
	if response.StatusCode == 200 {
		content, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println("Status Code : ", response.StatusCode)
		fmt.Println("Content Length : ", response.ContentLength)
		fmt.Println(string(content))

		// Strings builder has it's buffer of it own where all the data about any entity is stored once `Write(arg)` is used to fill the buffer then we can use `String()` method to dump all the accumulated value
		var responseString strings.Builder
		// write appends the content of the argument into (b *Builder) 's  buffer
		fmt.Println(responseString.Write(content))
		fmt.Println(responseString.String())
	}

}

// POSTFORM request takes a
func PerformPostFormRequest(myUrl string) {
	myUrl = myUrl + "/postform"
	data := url.Values{}
	data.Add("FirstName", "Vandit")
	data.Add("LastName", "Singh")
	data.Add("Singer", "Nanku")

	// checkout the http.PostForm func
	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
