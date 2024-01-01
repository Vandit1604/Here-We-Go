package main

import "fmt"

// visit runes later
// 🟠 String are just slice of bytes in golang not character array but runes. For more info go here ~> https://go.dev/blog/strings
func main() {
	// hello in thai language
	const s = "สวัสดี"
	fmt.Printf("Length of S:%v\n", len(s))

}
