package main

import "fmt"

func main() {
	a := 58
	var ptr *int
	ptr = &a

	fmt.Println("Value of a:", a) // Output: 58
	fmt.Println("Memory address of a:", &a)
	fmt.Println("Pointer ptr holds the address of a:", ptr)
	fmt.Println("Value at the address pointed to by ptr:", *ptr)
}
