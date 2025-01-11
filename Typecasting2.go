package main

import "fmt"

func main() {
	var floatVal float64 = 42.75
	var intVal int = int(floatVal) // Type casting from float64 to int

	fmt.Printf("Float: %.2f, Integer: %d\n", floatVal, intVal)
}
