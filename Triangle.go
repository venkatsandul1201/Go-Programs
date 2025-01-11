package main

import "fmt"

func main() {
	var rows int
	fmt.Print("Enter the number of rows for Pascal's Triangle: ")
	fmt.Scan(&rows)
	var pascalTriangle [][]int
	for i := 0; i < rows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = pascalTriangle[i-1][j-1] + pascalTriangle[i-1][j]
		}
		pascalTriangle = append(pascalTriangle, row)
	}

	fmt.Println("Pascal's Triangle:")
	for i := 0; i < rows; i++ {

		for j := 0; j < rows-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print(pascalTriangle[i][j], " ")
		}
		fmt.Println()
	}
}
