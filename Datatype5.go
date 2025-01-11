package main

import "fmt"

func main() {
	// Array
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)

	// Slice
	var numbersSlice []int = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", numbersSlice)
}
