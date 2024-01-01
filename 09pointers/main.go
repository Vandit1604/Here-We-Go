package main

// 🍏 A pointer = address will get you the direct access to that entity in golang; A POINTERS NEEDS TO POINT AND IF YOU GIVE IT AN ADDRESS YOU'RE BASICALLY TELLING IT WHERE TO POINT.

import (
	"fmt"
)

func valueZero(value int) {
	value = 0
}
func valueZeroPtr(ptr *int) {
	*ptr = 0
}
func main() {
	i := 1
	fmt.Println("Initially:", i)
	valueZero(i)
	fmt.Println("After Zero Value:", i)
	valueZeroPtr(&i)
	fmt.Println("After Zero Ptr:", i)
}
