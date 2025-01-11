package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world"
	fmt.Println(strings.ReplaceAll(s, "world", "Go"))
}
