package main

import "fmt"

func main() {
	var num1, num2 float64
	var choice int
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)
	fmt.Println("\nChoose an arithmetic operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Print("\nEnter your choice (1/2/3/4): ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Printf("\n%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case 2:
		fmt.Printf("\n%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case 3:
		fmt.Printf("\n%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case 4:
		if num2 != 0 {
			fmt.Printf("\n%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		} else {
			fmt.Println("\nError: Division by zero is not allowed.")
		}
	default:
		fmt.Println("\nInvalid choice. Please choose a valid operation.")
	}
}
