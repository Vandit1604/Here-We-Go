package main

import (
	"fmt"
	"net/http"
)

func main() {
	// This will only print World 5 times because as soon as the `greeter("World")` function will execute because the program exits and doesn't wait for the goroutine thread greeter("Hello") function
	// go greeter("Hello")
	// greeter("World")

	weblist := []string{"https://google.com", "https://youtube.com", "https://github.com", "https://hashnode.com", "https://elixircommunity.live"}
	for _, site := range weblist {
		getStatusCode(site)
	}
}

// func greeter(str string) {
// 	for i := 0; i < 6; i++ {
// 		fmt.Println(str)
// 	}
// }

// this will be a very slow response and we can use concurrency here for fast results
func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status code is %d for endpoint %s\n", res.StatusCode, endpoint)
}
