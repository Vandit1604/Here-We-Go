package main

import "fmt"

func main() {
	fmt.Println("Learning about Loops in go")

	// go only has for loop
	// single condition
	i := 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// normal for loop
	for i := 4; i < 10; i++ {
		fmt.Println(i)
	}

	// while loop
	for {
		i++
		if i == 15 {
			continue
		}
		if i == 18 {
			break // breaks out of the loop
		}
		fmt.Println("Hi", i)
	}

	// while
}
