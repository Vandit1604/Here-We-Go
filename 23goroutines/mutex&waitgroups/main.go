package main

import (
	"fmt"
	"net/http"
	"sync"
)

// WE WILL USE WAITGROUPS HERE

// VANDIT IN FUTURE YOU USED WAITGROUPS AND IT WAS VERY FAST 🚀

var wg sync.WaitGroup

// Basically lock a resource when one goroutine is accessing it. What I learnt in Operating systems will be handy here
var mutex sync.Mutex
var signal = []string{"test"}

// WaitGroups have the role of telling the program to WAIT until wg recieves a done signal. It also tracks the number of goroutines that are fired and we should wait for each of them.
func main() {
	weblist := []string{"https://google.com", "https://youtube.com", "https://github.com", "https://hashnode.com", "https://elixircommunity.live"}
	for _, site := range weblist {
		go getStatusCode(site)
		// add that 1 goroutine is fired right now
		// increase waitGroup counter by 1
		wg.Add(1)
	}

	// 	waits until waitGroup counter is zero
	wg.Wait()
	fmt.Println(signal)
}

func getStatusCode(endpoint string) {
	// decrease waitGroup counter value by 1
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status code is %d for endpoint %s\n", res.StatusCode, endpoint)
	mutex.Lock()
	signal = append(signal, endpoint)
	mutex.Unlock()
}
