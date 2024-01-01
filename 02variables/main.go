package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("We'll be learning about variables, but first some values")

	// string
	fmt.Println("This is a string")

	// int
	fmt.Println(12)

	// calculations
	fmt.Println("12/4:", 12/4)

	// booleans
	fmt.Println(true, false, true && false)

	// how to define a variable in golang
	// In go you define the type declaring it's name weird right ?
	var number int = 12
	var name string = "Vandit"
	var bool bool = true
	var float float32 = 12.33 / 3
	fmt.Println(number, name, bool, float)

	// there's also an awesome thing in golang called *walrus operator*
	// basically if you don't know the type of what's coming from a package(this happens when using different packages that you have not written) go will automatically infer it's type and use it.
	// looks like a walrus right ? (:=) when using this you don't have to right *var* go will infer everything. The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.
	some := 12 / 34
	fmt.Println("hi", some)

	// variables that are not initialized with a value are zero-valued.
	var zero int     // 0
	var nulll string // empty space
	var mid float32  // 0
	fmt.Println(zero)
	fmt.Println(nulll)
	fmt.Println(mid)

	// constants
	const n = 10000
	var iota = complex(3, 10)
	fmt.Println(iota)

	// you can use the reflect package to find the type of a variable
	fmt.Println(reflect.TypeOf(iota))
	fmt.Println(reflect.TypeOf(n))
	fmt.Println(n / iota)
}
