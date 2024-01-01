package main

import "fmt"

// if a normal statement contains `defer` imagine it will be added in a stack and the stack will pop until its empty only when it gets to the closing bracket of main function

// Q. Why is Defer used?
// A. When you want to do some task and then clean up or close the stream or entity access you use defer so it can be automatically closed when you get to the end of last closing bracket

// output should be : 1 4 3 2 1 0 4 3 2
func main() {
	fmt.Println("1")

	// multiple defer statements follow the same put it stack approach. THAT EXAMPLE IS FOR BETTER UNDERSTANDING
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")

	// normally it should print
	// 1 2 3 4

	// when we use defer the output will start from bottom to top
	// 1 4 3 2
	Defer()
}

func Defer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
		// 43210
	}
}
