package main

import "fmt"

func main() {
	str := "Hello, World!"
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range str {
		if contains(vowels, char) {
			count++
		}
	}
	fmt.Println("Number of vowels:", count)
}
func contains(vowels string, char rune) bool {
	for _, v := range vowels {
		if v == char {
			return true
		}
	}
	return false
}
