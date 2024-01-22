package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Starting server....")

	mux := http.NewServeMux()
	mux.HandleFunc("/greet/", greet)
	mux.HandleFunc("/calculator/", calculator)
	mux.HandleFunc("/", defaultFunc)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/ root request")
	fmt.Fprintf(w, "Hello World")
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/greet request")
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(w, "Hello Ji")
	} else {
		fmt.Fprintf(w, "Hello "+name)
	}
}

func calculator(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/calculator request")
	op1 := r.URL.Query().Get("op1")
	op := r.URL.Query().Get("op")
	op2 := r.URL.Query().Get("op2")

	if op1 == "" || op2 == "" || op == "" {
		fmt.Fprintf(w, "Please provide all the parameters")
	} else {
		fmt.Fprintf(w, "Result is %#v", calculate(op1, op, op2))
	}
}

func calculate(op1, op, op2 string) int {
	int_op1, err := strconv.Atoi(op1)
	int_op2, err := strconv.Atoi(op2)
	if err != nil {
		log.Fatal(err)
	}
	switch op {
	case "add":
		return int_op1 + int_op2
	case "-":
		return int_op1 - int_op2
	case "*":
		return int_op1 * int_op2
	case "/":
		return int_op1 / int_op2
	default:
		return 0
	}
}
