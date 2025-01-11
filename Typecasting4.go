package main

import "fmt"

func main() {
	var intVal int = 123
	var strVal string = fmt.Sprintf("%d", intVal) // Type casting from int to string

	fmt.Printf("Integer: %d, String: %s\n", intVal, strVal)
}
