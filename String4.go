package main

import (
	"fmt"
	"strings"
)

func main() {
	// Input string
	input := "hello, world!"

	// Convert to uppercase
	uppercase := strings.ToUpper(input)

	// Print the result
	fmt.Println("Original String:", input)
	fmt.Println("Uppercase String:", uppercase)
}