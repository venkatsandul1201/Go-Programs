package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	fmt.Println("Multiplication table of", number)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}
