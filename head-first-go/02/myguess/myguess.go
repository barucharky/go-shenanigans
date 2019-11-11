package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// -- -----------------------------------
	// set random number
	var seconds int64 = time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I chose a random number 1 <> 10.")
	fmt.Println("Can you guess it?")
	// -- -----------------------------------

	// -- -----------------------------------
	// Set up the reader for keyboard input:
	reader := bufio.NewReader(os.Stdin)

	// -- -----------------------------------
	// play the game
	var success bool = false

	// Loop time
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Enter your guess: ")

		// -- -----------------------------------
		// Get user input
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}
		// -- -----------------------------------

		input = strings.TrimSpace(input)

		// Convert string to integer (Atoi)
		guess, err := strconv.Atoi(input)

		if err != nil {
			guess--
			fmt.Println("Invalid guess. Try again.")
			continue
		}

		// -- -----------------------------------
		// Check guess:
		if guess < target {
			fmt.Println("Too LOW.")
		} else if guess > target {
			fmt.Println("Too HIGH")
		} else {
			success = true
			fmt.Println("You guessed it!")
			break
		}
	}

	if !success {
		fmt.Println("You lose. It was", target)
	}
}
