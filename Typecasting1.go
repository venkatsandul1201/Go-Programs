package main

import "fmt"

func main() {
	var intVal int = 42
	var floatVal float64 = float64(intVal) // Type casting from int to float64

	fmt.Printf("Integer: %d, Float: %.2f\n", intVal, floatVal)
}
