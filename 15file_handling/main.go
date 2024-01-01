package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	content := "Hi I'll be in a file in a bit"

	// Returns a pointer to the file
	file, err := os.Create("./file.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is ", length)
	// this defer statement will be executed at the very last
	defer file.Close()
	ReadFile("./file.txt")

	// Reading a directory
	directoryname, err := os.ReadDir("../../golang-here-we-go")
	fmt.Printf("%+v", directoryname)
	for _, name := range directoryname {
		fmt.Println(name)
	}

	// let's create a dir for the next module
	// os.Mkdir("../16web-request/", 0777)

	mode := int(0777)
	os.Mkdir("../16web-request/", os.FileMode(mode))

	// let's also create a file inside the 16web-request dir
	newfile, err := os.Create("../16web-request/main.go")
	if err != nil {
		panic(err)
	}

	basic := "package main\n func main(){}\n"
	io.WriteString(newfile, basic)
}

func ReadFile(filename string) {
	readContent, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(readContent))
}
