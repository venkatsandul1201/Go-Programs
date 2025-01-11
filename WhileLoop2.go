package main

import "fmt"

func main() {
	sum := 0
	i := 1
	for i <= 100 { // Simulating a while loop
		sum += i
		i++
	}
	fmt.Println("Sum of numbers from 1 to 100 is:", sum)
}
