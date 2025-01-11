package main

import "fmt"

func main() {
	var char string
	fmt.Print("Enter a character: ")
	fmt.Scanln(&char)

	switch char {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		fmt.Println("The character is a vowel.")
	case "b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z",
		"B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "X", "Y", "Z":
		fmt.Println("The character is a consonant.")
	default:
		fmt.Println("The character is a special symbol or number.")
	}
}
