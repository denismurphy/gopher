package game

import (
	"fmt"
	"math/rand"
	"time"
)

// GuessTheNumber represents a number guessing game
type GuessTheNumber struct{}

// Guess starts a number guessing game
func (g *GuessTheNumber) Guess() {
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100) + 1
	fmt.Println("Guess the number between 1 and 100!")

	for {
		var guess int
		fmt.Print("Enter your guess: ")

		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if guess < secretNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > secretNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Congratulations! You guessed the number!")
			break
		}
	}
}

// DemonstrateGuessGame starts a demonstration of the guessing game
func DemonstrateGuessGame() {
	fmt.Println("Starting the Guess The Number game...")
	guessTheNumber := GuessTheNumber{}
	guessTheNumber.Guess()
}
