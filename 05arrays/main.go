package main

import "fmt"

func main() {
	// Let's create arrays today and have some fun with them.
	// this is how you create arrays

	var a [5]int
	fmt.Println(a)
	// initially the array is zero valued in case of integer it's 0

	// initializing with values
	// need to practice this one
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// array works similar to arrays in diff. language

	// getting and setting values in of an existing array

	// setting
	b[0] = 9
	fmt.Println(b)

	// getting
	fmt.Println(b[0])

	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println("Array Incoming", arr1)

	// some sneak peak of 2D array
	var we [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			we[i][j] = i * j
		}
	}

	fmt.Println(we)
}
