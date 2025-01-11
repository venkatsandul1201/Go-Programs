package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&n)

	a, b := 0, 1

	fmt.Println("Fibonacci Series:")
	for i := 0; i < n; i++ {

		fmt.Print(a, " ")
		nextTerm := a + b
		a = b
		b = nextTerm
	}

	fmt.Println()
}
