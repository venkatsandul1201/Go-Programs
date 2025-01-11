package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)
	sum := 0
	i := 1
	for {
		sum += i
		i++
		if i > n {
			break
		}
	}
	fmt.Println("Sum of numbers from 1 to", n, "is:", sum)
}
