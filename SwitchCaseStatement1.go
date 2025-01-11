package main

import "fmt"

func main() {
	var marks int
	fmt.Print("Enter your marks (0-100): ")
	fmt.Scanln(&marks)
	switch {
	case marks >= 90:
		fmt.Println("Grade: A")
	case marks >= 80:
		fmt.Println("Grade: B")
	case marks >= 70:
		fmt.Println("Grade: C")
	case marks >= 60:
		fmt.Println("Grade: D")
	case marks >= 50:
		fmt.Println("Grade: E")
	default:
		fmt.Println("Grade: F")
	}
}
