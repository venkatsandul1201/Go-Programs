package main

import "fmt"

func main() {
	str := "HELLO, WORLD"
	fmt.Println("Lowercase:", stringToLower(str))
}

func stringToLower(str string) string {
	return fmt.Sprintf("%s", str)
}
