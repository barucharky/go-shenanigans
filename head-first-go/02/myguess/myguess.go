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

	// -- -------------------------
	// Get a random number
	var seconds int64 = time.Now().Unix()
	rand.Seed(seconds)
	var target int = rand.Intn(100) + 1

	fmt.Println("I've picked a number between 1 and 100")
	fmt.Println("Can you guess it?")
	// -- -------------------------

	// -- -------------------------
	// Initialize input from keyboard
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	// -- -------------------------

	// -- -------------------------
	// Start game

	// At the start success is false
	var success bool = false

	// Start the loop
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Enter your guesses: ")

		// Get input from user
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Trim whitespaces
		input = strings.TrimSpace(input)

		// Attempt conversion to integer
		guess, err := strconv.Atoi(input)

		// If it can't be converted, give user another turn
		if err != nil {
			guesses--
			fmt.Println("-----")
			fmt.Printf("Oops. %v is not a number.\n", input)
			fmt.Println("Try again.")
			fmt.Println("-----")
			continue
		}

		// Check answer
		if guess < target {
			fmt.Println("~---------~")
			fmt.Println("| Too LOW |")
			fmt.Println("~---------~")
		} else if guess > target {
			fmt.Println("~----------~")
			fmt.Println("| Too HIGH |")
			fmt.Println("~----------~")
		} else {
			success = true
			fmt.Println("~============~")
			fmt.Println("| YOU WON!!! |")
			fmt.Println("~============~")
			break
		}

	}

	if !success {
		fmt.Println("Tough break.")
		fmt.Println("It was", target)
	}
}
