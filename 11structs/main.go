package main

import "fmt"

type Bottle struct {
	Color           string
	QuantityInLiter int
	Filled          bool
	Contains        string
}

// func to fill bottle with milk

// the pointer struct doesnt operate on the copy of the struct
func (b *Bottle) FillUpMilk() {
	if b.Filled == false || b.Contains != "Milk" {
		b.Filled = true
		b.Contains = "Milk"
	}
}

// uncomment this code and check the output for	fmt.Println("After milk fillup", hamilton)
// func (b Bottle) FillUpMilk() {
// 	if b.Filled == false || b.Contains != "Milk" {
// 		b.Filled = true
// 		b.Contains = "Milk"
// 	}
// }

// defining that it will be returning a pointer to the Bottle struct
func defaultBottle(color string) *Bottle {
	p := Bottle{Color: color}
	p.Contains = "Water"
	p.Filled = false
	p.QuantityInLiter = 1
	return &p // returns address of the struct we created inside the function so the struct initialised inside can be accessed outside the function too via the pointer
}

func main() {
	fmt.Println("Let's learn structs")
	fmt.Println("---")

	milton := Bottle{"red", 1, true, "Water"}
	fmt.Printf("%+v\n", milton)
	ptrMilton := &milton
	// fmt.Println(ptrMilton.Contains)
	ptrMilton.Contains = "Milk"
	// fmt.Printf("%+v\n", milton)
	// fmt.Println(ptrMilton.Contains)

	// creating a struct using a function
	hamilton := defaultBottle("silver")
	hamilton.FillUpMilk()
	milton.FillUpMilk()
	fmt.Println("After milk fillup", milton)
	fmt.Println("After milk fillup", hamilton)

	// if you want some collection of fields that will only be used once, you can create anon nameless structs
	cat := struct {
		IsCute bool
		Name   string
	}{
		true,
		"Gajju",
	}
	fmt.Printf("%+v\n", cat)
}
