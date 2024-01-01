package main

import (
	"fmt"
	"time"
)

func main() {
	// normal if else like what we have in other languages
	var i int = 0
	if i > 5 {
		fmt.Println("i is greater than 5")
	} else if i == 0 {
		fmt.Println("i is zero")
	} else {
		fmt.Println("i is smaller than 5")
	}
	fmt.Println("hi", time.Now().Weekday())

	//switch case
	switch time.Now().Year()/400 == 0 {
	case true:
		fmt.Println("It's leap year")
	default:
		fmt.Println("It's not leap year")
	}

	// let's try to make a clock
	for {
		t := time.Now().Second()
		t1 := time.Now().Second()
		if t1-t == 1 {
			fmt.Println(t1)
		}
	}
}
