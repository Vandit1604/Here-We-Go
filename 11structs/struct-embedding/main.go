package main

import (
	"fmt"
)

type Gemstone struct {
	Color, Size string
}

func (g *Gemstone) Polish() {
	g.Color = "polished " + g.Color
}

type Wall struct {
	Gemstone
	Length, Width int
}

type person struct {
	name string
	id   uint
	kewl bool
}

func (p *person) makeKewl() {
	p.kewl = true
}

func main() {
	wall1 := Wall{

		// The name of the struct should be here just like this name: name{
		// 	field: value,
		// 	field: value,
		// }

		Gemstone: Gemstone{
			Color: "red",
			Size:  "Large",
		},
		Length: 12,
		Width:  12,
	}
	fmt.Println(wall1)
	// We can create functions for a struct and implement them easily on container structs.
	// Since container embeds base, the methods of base also become methods of a container. Here we invoke a method that was embedded from base directly on co.
	// where container ~> struct that got base embedded into
	// base ~> struct that is embedded

	wall1.Polish()

	// embedded struct fields can be directly accessed, see below
	fmt.Println(wall1.Color)

	// we can also use the full path
	fmt.Println("Size", wall1.Gemstone.Size)

	vandit := person{"Vandit", 1, false}
	fmt.Printf("%+v\n", vandit)
	vandit.makeKewl()
	fmt.Printf("%+v\n", vandit)
}
