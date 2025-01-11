package main

import "fmt"

func main() {
	var a, b, z int
	z = 1
	fmt.Printf("\n Enter base and power:")
	fmt.Scanf("%d %d", &a, &b)
	for i := 1; i <= b; i++ {
		z = z * a
	}
	fmt.Printf("%d Raise to %d=%d\n", a, b, z)
}
