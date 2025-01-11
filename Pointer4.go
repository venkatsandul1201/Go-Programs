package main

import "fmt"

func modifyArray(arr *[3]int) {
	arr[0] = 100 

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println("Before modification:", arr) 

	modifyArray(&arr)                       
	fmt.Println("After modification:", arr) 
}
