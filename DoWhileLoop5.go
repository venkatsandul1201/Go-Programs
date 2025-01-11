package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(10) + 1
	var guess int
	for {
		fmt.Print("Guess the number (between 1 and 10): ")
		fmt.Scanln(&guess)
		if guess == secretNumber {
			fmt.Println("Congratulations! You guessed the right number.")
			break
		} else {
			fmt.Println("Incorrect guess. Try again!")
		}
	}
}
