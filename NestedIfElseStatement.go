package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	if num > 0 {
		if num%2 == 0 {
			fmt.Println("The number is positive and even.")
		} else {
			fmt.Println("The number is positive and odd.")
		}
	} else if num < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}
}
