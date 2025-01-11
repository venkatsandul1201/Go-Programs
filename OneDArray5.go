package main

import "fmt"

func main() {

	arr := [4]float64{10.5, 20.5, 30.5, 40.5}

	sum := 0.0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	average := sum / float64(len(arr))

	fmt.Printf("The average of the array elements is: %.2f\n", average)
}
