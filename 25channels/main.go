package main

import (
	"fmt"
	"sync"
)

// Channels are like telephone lines that can't be used until there's someone at both the ends one sending something and one recieving it.
func main() {
	fmt.Println("Let's Learn Channels")
	var wg sync.WaitGroup

	// this is how to create a channel. This is a more generalized way that's commonly followed by people
	// you can also make buffered channels which are used when more than one value will be pass on a channel and only 1 reciever will listen to those values
	myChannel := make(chan int)

	wg.Add(2)

	// NOTE: You can listen to a channel even if no one is sending anything to it. but you can't send data if someone is not recieving data. It will give 0
	go func(ch chan int) {
		fmt.Println(<-ch)
		wg.Done()
	}(myChannel)

	go func(ch chan int) {
		ch <- 5
		close(ch)
		wg.Done()
	}(myChannel)

	wg.Wait()
}
