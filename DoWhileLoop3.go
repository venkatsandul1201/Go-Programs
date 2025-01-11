package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	factorial := 1
	i := 1
	for {
		factorial *= i
		i++
		if i > num {
			break
		}
	}
	fmt.Println("Factorial of", num, "is:", factorial)
}
