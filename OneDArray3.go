package main

import "fmt"

func main() {
	arr := [6]int{10, 25, 15, 30, 5, 50}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Printf("The maximum element in the array is: %d\n", max)
}
