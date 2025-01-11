package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if num > 0 {
		fmt.Println("The number is positive.")
	}
	if num < 0 {
		fmt.Println("The number is negative.")
	}
	if num == 0 {
		fmt.Println("The number is zero.")
	}
}
