package main

import "fmt"

func concatenateStrings(str1, str2 *string) string {
	return *str1 + *str2
}
func main() {
	str1 := "Hello, "
	str2 := "World!"

	ptr1 := &str1
	ptr2 := &str2
	result := concatenateStrings(ptr1, ptr2)
	fmt.Println("Concatenated String:", result)
}
