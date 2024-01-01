package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

// const myurl string = "https://example.com:3000/vandit?cool=true&swag=true"

func main() {
	fmt.Println("Let's learn about URLs in golang")

	// parsing the URL
	// result, _ := url.Parse(myurl)
	// fmt.Println(result.Host)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	// values := result.Query()
	// fmt.Println(values)

	// Building URLs
	ytSearchURL := &url.URL{
		Scheme:   "https",
		Host:     "youtube.com",
		Path:     "results",
		RawQuery: "search_query=",
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What do you wanna search for")
	searchquery, _ := reader.ReadString('\n')
	searchquery = strings.ReplaceAll(searchquery, " ", "+")
	fmt.Printf("%v%v\n", ytSearchURL.String(), searchquery)
}
