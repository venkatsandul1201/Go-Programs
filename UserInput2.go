package main

import "fmt"

func main() {
	var name string
	var age int
	var height float64
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Print("Enter your height (in meters): ")
	fmt.Scanln(&height)
	fmt.Printf("Name: %s\nAge: %d\nHeight: %.2f meters\n", name, age, height)
}
