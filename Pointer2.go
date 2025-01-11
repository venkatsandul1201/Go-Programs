package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func main() {
	var funcPtr func(int, int) int

	funcPtr = add
	fmt.Println("Addition:", funcPtr(10, 5))

	funcPtr = subtract
	fmt.Println("Subtraction:", funcPtr(10, 5))
}
