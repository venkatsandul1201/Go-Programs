package main

import "fmt"

func modifyValue(val *int) {
	*val = *val + 10
}

func main() {
	a := 5
	fmt.Println("Before modification:", a)

	modifyValue(&a)
	fmt.Println("After modification:", a)
}
