package main

import (
	"fmt"
)

func main() {
	// slices are arrays that don't require the size when declared

	var slice []string
	fmt.Println("This is an uninitialized slice", slice, slice == nil, len(slice) == 0)

	s := make([]int, 3)
	for i := 0; i < len(s); i++ {
		fmt.Println(i, "=>", s[i])
	}
	s[0] = 1
	fmt.Println(s)

	c := make([]int, 3)
	copy(c, s)
	fmt.Println(c)

	slicee := []int{1, 2, 3, 4, 5}
	fmt.Println(slicee)

	// slicing the slices
	l := slicee[2:4]
	fmt.Println(l)
	fmt.Println(slicee)

	// removing a value using the append method
	var friends = []string{"Tishika", "Stuti", "Shruti", "Vishal", "Bhanu"}
	friends = append(friends[:2], friends[3:]...)
	fmt.Println(friends)
}
