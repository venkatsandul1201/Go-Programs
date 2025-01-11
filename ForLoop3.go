package main

import "fmt"

func main() {
	var n, factorial int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	factorial = 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}

	fmt.Printf("Factorial of %d is: %d\n", n, factorial)
}
