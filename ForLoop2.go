package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Printf("Sum of first %d natural numbers is: %d\n", n, sum)
}
