package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var b byte = 'a'
	fmt.Println("Printing Byte:")

	fmt.Printf("Size: %d\nType: %s\nCharacter: %c\n", unsafe.Sizeof(b), reflect.TypeOf(b), b)

	r := '£'
	fmt.Println("\nPrinting Rune:")

	fmt.Printf("Size: %d\nType: %s\nUnicodeCodePoint: %U\nCharacter: %c\n", unsafe.Sizeof(r),
		reflect.TypeOf(r), r, r)

	s := "µ"
	fmt.Println("\nPrinting String:")
	fmt.Printf("Size: %d\nType: %s\nCharacter: %s\n", unsafe.Sizeof(s), reflect.TypeOf(s), s)
}
