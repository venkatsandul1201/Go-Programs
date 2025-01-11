package main

import "fmt"

func main() {
	var limit int
	fmt.Print("Enter the limit for the Fibonacci series: ")
	fmt.Scanln(&limit)
	a, b := 0, 1
	for {
		if a > limit {
			break
		}
		fmt.Println(a)
		a, b = b, a+b
	}
}
