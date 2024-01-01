package main

import "fmt"

// Interface will contains method signatures but not the method body just like abstract functions. The Drawback for interfaces is that you will need to have function body for each of the function even if a struct only needs one of the method to use with function body of its own.
type Traits interface {
	walk()
	fly()
}

// Create your structs
type Human struct {
	Name     string
	Activity string
}

// respective functions
func (h Human) walk() {

	fmt.Println("Human will Walk")
}

func (h Human) fly() {
	fmt.Println("Human will Not Fly")
}

type Bird struct {
	Name     string
	Activity string
}

// respective functions
func (b Bird) fly() {
	fmt.Println("Bird will Fly")
}

func (b Bird) walk() {

	fmt.Println("Bird will Walk")
}

// Create a function that will take the interface as input to reference the method signatures. the method body will be taken automatically to trigger the function
func Activity(a Traits) {
	a.walk()
	a.fly()
}
func main() {
	b := Bird{"sparrow", "nothing"}
	h := Human{"Vandit", "chill"}
	Activity(b)
	Activity(h)
}
