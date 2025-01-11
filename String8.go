package main

import "fmt"

func countOccurrences(s string, char rune) int {
	count := 0
	for _, v := range s {
		if v == char {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countOccurrences("hello world", 'o'))
}
