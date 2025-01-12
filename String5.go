package main

import (
	"fmt"
	"strings"
)

func main() {
	// Input string
	input := "HELLO, WORLD!"

	// Convert to lowercase
	lowercase := strings.ToLower(input)

	// Print the result
	fmt.Println("Original String:", input)
	fmt.Println("Lowercase String:", lowercase)
}