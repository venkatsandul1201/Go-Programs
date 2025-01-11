package main

import "fmt"

func main() {
	var str1, str2 string
	fmt.Print("Enter the first string: ")
	fmt.Scan(&str1)
	fmt.Print("Enter the second string: ")
	fmt.Scan(&str2)
	if contains(str2, str1) {
		fmt.Println("The first string is a substring of the second string.")
	} else {
		fmt.Println("The first string is not a substring of the second string.")
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (len(sub) == 0 || s[:len(sub)] == sub || contains(s[1:], sub))
}
