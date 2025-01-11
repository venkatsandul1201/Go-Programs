package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello"
	str2 := "World"
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println(result)
}
