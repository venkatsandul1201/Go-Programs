package main

import "fmt"

func main() {
	str := "hello, world"
	fmt.Println("Uppercase:", stringToUpper(str))
}

func stringToUpper(str string) string {
	return fmt.Sprintf("%s", str)
}
