package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<nb_max>")
		fmt.Println("Example:", os.Args[0], "42000")
		return
	}

	max, err := strconv.Atoi(os.Args[1])
	if err != nil || max <= 0 {
		fmt.Println("Invalid input. Please provide a positive integer.")
		return
	}

	numberToGuess, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		fmt.Println("Failed to generate a random number.")
		return
	}

	for {
		fmt.Print("Guess a number: ")
		var input string
		fmt.Scanln(&input)

		guessedNumber, err := strconv.Atoi(input)
		if err != nil || guessedNumber < 0 || guessedNumber >= max {
			fmt.Println("Invalid input. Please provide a number between 0 and", max-1)
			continue
		}

		formatedGuessedNumber := big.NewInt(int64(guessedNumber)) // same format as numberToGuess 

		switch formatedGuessedNumber.Cmp(numberToGuess) {
		case -1:
			fmt.Println("Too low!")
		case 1:
			fmt.Println("Too high!")
		case 0:
			fmt.Println("Congratulations! You guessed the number:", numberToGuess)
			return
		}
	}
}
