package main

import "fmt"

func main() {
	var ch string
	fmt.Print("Enter a character: ")
	fmt.Scanln(&ch)

	if ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u" ||
		ch == "A" || ch == "E" || ch == "I" || ch == "O" || ch == "U" {
		fmt.Println("The character is a vowel.")
	} else {
		fmt.Println("The character is a consonant.")
	}
}
