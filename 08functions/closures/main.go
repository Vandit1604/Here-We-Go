package main

import "fmt"

// 😼 ELI5 CLOSURES
func clock() func() int {
	// variable that will be bound to the enviorment inside this function called closure
	second := 0

	// actual function that will be returned, go to 🔵
	return func() int {
		second++
		return second
	}
}
func main() {

	// 🔵 Here what will happen is I'll assign this function to a variable and that variable will get the anonymous function which will triggered everytime that variable will be called, see below
	clock1 := clock()
	fmt.Println(clock1())
	fmt.Println(clock1())
	fmt.Println(clock1())
	fmt.Println(clock1())

	// but if i create another variable it will again start with 1
	clock2 := clock()
	fmt.Println(clock2())
}
