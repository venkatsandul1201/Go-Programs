package main

import "fmt"

func main() {
	str1 := "hello"
	str2 := "hello"
	if str1 == str2 {
		fmt.Println("The strings are equal")
	} else {
		fmt.Println("The strings are not equal")
	}
}
