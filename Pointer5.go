package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func updateAge(p *Person, newAge int) {
	p.Age = newAge

func main() {
	person := Person{Name: "Alice", Age: 30}
	fmt.Println("Before update:", person) 

	updateAge(&person, 35)               
	fmt.Println("After update:", person) 
}
