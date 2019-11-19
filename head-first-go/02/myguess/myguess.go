package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// -- --------------------------------
	// Pick a random number
	var seconds int64 = time.Now().Unix()
	rand.Seed(seconds)
	var target int = rand.Intn(100) + 1

	fmt.Println("I'm thinking of a number between 1 and 100")
	fmt.Println("Try to guess what it is.")

	// -- --------------------------------
	// Initialize input from keyboard
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	// -- --------------------------------
	// Play the game
	var success bool = false

	// Start the loop
	for guesses := 0; guesses < 10; guesses++ {

		if guesses == 0 {
			fmt.Println("Begin.")
			fmt.Println("-----")
		}

		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Enter your guesses: ")

		// get user input
		input, err := reader.ReadString('\n')

		// trim white spaces
		input = strings.TrimSpace(input)

		// Attempt conversion to int
		guess, err := strconv.Atoi(input)

		// If it fails, give another guess
		if err != nil {
			guesses--
			fmt.Printf("%s is not a number.\n", input)
			fmt.Println("Guess again")
			continue
		}

		// Check guess
		if guess < target {
			fmt.Println("Too LOW")
		} else if guess > target {
			fmt.Println("Too HIGH")
		} else {
			success = true
			fmt.Println("You got it!")
			break
		}
	}

	if !success {
		fmt.Println("You lose.")
		fmt.Printf("%d was the right answer.\n", target)
	}
}
