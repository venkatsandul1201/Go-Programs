package main

import "fmt"

func main() {
	str := "madam"
	isPalindrome := true

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("The string is a palindrome")
	} else {
		fmt.Println("The string is not a palindrome")
	}
}
