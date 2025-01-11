package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "helloworldgo"
	words := strings.Split(sentence, " ")
	fmt.Println(words)
}
