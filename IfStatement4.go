package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if num%5 == 0 {
		fmt.Println("The number is a multiple of 5.")
	}
}
