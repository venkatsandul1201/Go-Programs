package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("After swapping:")
	fmt.Println("First number:", a)
	fmt.Println("Second number:", b)
}
