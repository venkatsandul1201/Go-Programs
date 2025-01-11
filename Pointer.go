package main

import "fmt"

func main() {

	var num int = 58
	var ptr1 *int = &num
	var ptr2 **int = &ptr1
	fmt.Println("Value of num:", num)
	fmt.Println("Value pointed to by ptr1:", *ptr1)
	fmt.Println("Value pointed to by ptr2:", **ptr2)
	**ptr2 = 100
	fmt.Println("Modified value of num:", num)
}
