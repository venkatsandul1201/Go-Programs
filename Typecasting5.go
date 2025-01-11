package main

import "fmt"

func main() {
	var f32 float32 = 3.14
	var f64 float64 = float64(f32)    // Converting float32 to float64
	var f32New float32 = float32(f64) // Converting float64 back to float32

	fmt.Printf("Original float32: %.2f, Converted to float64: %.2f\n", f32, f64)
	fmt.Printf("Converted back to float32: %.2f\n", f32New)
}
