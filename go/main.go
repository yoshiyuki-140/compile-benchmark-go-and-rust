package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	// print message
	println("Guess the number!")

	// generate random integer
	secret_number := rand.Intn(100)
	println(secret_number)

	// define variable
	var guess string
	var guess_int int

	for {
		println("Please input your guess.")

		// read line
		_, err := fmt.Scan(&guess)
		if err != nil {
			println("Failed to read line")
			break
		}

		// convert to int
		guess_int, err = strconv.Atoi(guess)
		if err != nil {
			println("Failed to convert guess variable to `int` type variable from `string` type variable")
			break
		}
		println(fmt.Sprintf("You guessed: %d", guess_int))

		// condition judge
		if secret_number == guess_int {
			println("You win!")
			break
		} else if secret_number < guess_int {
			println("Too big!")
		} else {
			println("Too small!")
		}
	}
}
