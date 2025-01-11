package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "123"
	var num int
	var err error
	// Converting string to int using strconv.Atoi()
	num, err = strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("String: %s, Integer: %d\n", str, num)
	}
}
