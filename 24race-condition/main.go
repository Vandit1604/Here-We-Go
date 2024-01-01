package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Let's create a race condition")

	var score []string
	var wg sync.WaitGroup

	// we can checking race conditions if we run the code with commands : go run --race main.go

	// we can use mutex lock to solve the race condition
	var mut sync.Mutex

	// always lock the resource when you're performing something on a resource
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("1")
		mut.Lock()
		score = append(score, "1")
		mut.Unlock()
		wg.Done()
	}(&wg, &mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("2")
		mut.Lock()
		score = append(score, "2")
		mut.Unlock()
		wg.Done()
	}(&wg, &mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("3")
		mut.Lock()
		score = append(score, "3")
		mut.Unlock()
		wg.Done()
	}(&wg, &mut)

	wg.Wait()
	fmt.Println(score)
}
