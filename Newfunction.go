package main

import "fmt"

func main() {
	ptr := new(int)
	fmt.Println("Initial value of ptr:", *ptr)
	*ptr = 58
	fmt.Println("Updated value of ptr:", *ptr)
	strPtr := new(string)
	fmt.Println("Initial value of strPtr:", *strPtr)
	*strPtr = "Hello, Go!"
	fmt.Println("Updated value of strPtr:", *strPtr)
}
