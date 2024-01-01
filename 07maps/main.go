package main

import (
	"fmt"
)

func main() {
	// 🔺 Maps are just like dictionaries

	// 🔺 Empty maps dont print anything

	// empty_map := make(map[int]int)
	// empty_map_string := make(map[string]string)
	// fmt.Println(empty_map, empty_map_string)

	languages := make(map[string]string)

	// 🔺 INITIALISING IN A SINGLE LINE
	var food = map[int]string{0: "HI", 1: "GOLANG"}
	fmt.Println(food)

	// 🔺 Adding a key:value pair
	languages["py"] = "Python"
	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["vi"] = "Vimscript"
	fmt.Println(languages)

	// 🔺 Reading value using a key
	fmt.Println(languages["py"])

	// 🔺 When key is not available it returns zero value for that datatype
	fmt.Println("Key doesn't exist that's why you get: ", languages["ty"])

	// 🔺 UPDATING THE KEY VALUE PAIR; you can update the value referencing a key

	languages["py"] = "Pycharm"
	fmt.Println(languages)

	// 🔺 len() function returns length of map
	fmt.Println("Length of the map is", len(languages))

	// clear() and delete()

	// empties the map
	// clear(languages)

	fmt.Println(languages)

	// DELETING A VALUE USING A KEY
	delete(languages, "rb")
	fmt.Println(languages)

	// HOW TO CHECK IF A VALUE EXISTS IN A MAP OR NOT
	languages["hbs"] = "Handlebars"

	_, exists := languages["hbs"]
	fmt.Println(exists)

	_, exists = languages["bs"]
	fmt.Println(exists)

	// RANGE | LOOPING THROUGH A MAP
	// WORKS LIKE WHILE TRUE LOOP

	// for {
	// 	fmt.Println("VANDIT")
	// }

	for _, value := range languages {
		fmt.Printf("Value: %v\n", value)
		// fmt.Printf("Key: %v, Value:%v\n", key, value)
	}

	// you can also iterate over the strings
	for i, unicode := range "go" {
		fmt.Println(i, unicode)
	}
}
