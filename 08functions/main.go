package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func fib() (int, int) {
	return 0, 1
}

func sum(nums ...int) {
	sum := 0
	// for range gives index and value in same order
	for _, nums := range nums {
		sum += nums
	}
	// fmt.Println(nums)
	fmt.Println(sum)
}

// well done vandit
func ConcatAll(word ...string) {
	concatenated := ""
	for _, i := range word {
		concatenated += i
	}
	fmt.Println(concatenated)
}
func main() {
	res := plus(1, 2)
	fmt.Println(res)
	res = plusPlus(1, 2, 3)
	fmt.Println(res)
	fmt.Println(fib())
	_, a := fib()
	fmt.Println(a)

	// Variadic Functions
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// Variadic Function for Concatenate all words
	words := []string{"hi", "vandit", "this", "will", "be", "concatenated"}
	ConcatAll(words...)
}
