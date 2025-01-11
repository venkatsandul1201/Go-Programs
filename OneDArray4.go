package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	fmt.Println("Reversed array:", arr)
}
