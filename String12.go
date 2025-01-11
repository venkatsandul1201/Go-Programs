package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world, hello go"
	fmt.Println(strings.LastIndex(s, "hello"))
}
