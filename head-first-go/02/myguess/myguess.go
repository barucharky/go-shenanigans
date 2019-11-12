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

	// -- ------------------------
	// Pick a random number
	var seconds int64 = time.Now().Unix()
	rand.Seed(seconds)
	var target int = rand.Intn(100) - 1
	fmt.Println("~----------------------------------------~")
	fmt.Println("| I've picked a number between 1 and 100 |")
	fmt.Println("| See if you can guess it                |")
	fmt.Println("~----------------------------------------~")
	// -- ------------------------

	// -- ------------------------
	// Initialize keyboard input
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	// -- ------------------------

	// -- ------------------------
	// Play the game
	var success bool = false

	//Start the loop
	for guesses := 0; guesses < 10; guesses++ {

		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Enter your guess: ")

		// Get user input
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Trim string
		input = strings.TrimSpace(input)

		// Convert to integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			guesses--
			fmt.Println("Whoa, that's not a number")
			fmt.Println("Try again")
			continue
		}

		// Check guess
		if guess < target {
			fmt.Println("Too LOW")
		} else if guess > target {
			fmt.Println("Too HIGH")
		} else {
			success = true
			fmt.Println("You win!!!")
			break
		}
	}

	if !success {
		fmt.Println("You're out of guesses.")
		fmt.Println("It was", target)
	}
}
