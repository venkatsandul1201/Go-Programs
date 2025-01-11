package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	factorial := 1
	i := 1
	for i <= num {
		factorial *= i
		i++
	}
	fmt.Println("Factorial of", num, "is:", factorial)
}
