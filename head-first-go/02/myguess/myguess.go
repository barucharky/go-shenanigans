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

	// Get a random number
	var seconds int64 = time.Now().Unix()
	rand.Seed(seconds)
	var target int = rand.Intn(100) + 1

	fmt.Println("I've picked a number between 1 and 100")
	fmt.Println("See if you can guess what it is.")

	// Initialize input from keyboard
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	// Play the game
	var success bool = false

	// Start the loop
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Enter your guess: ")

		var input string
		var err error

		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		// Convert to int or give another turn
		var guess int
		guess, err = strconv.Atoi(input)
		if err != nil {
			guesses--
			fmt.Println("-----")
			fmt.Println("Sorry,", input, "is not a number.")
			fmt.Println("Try again.")
			fmt.Println("-----")
			continue
		}

		// Check answer
		if guess < target {
			fmt.Println("-----")
			fmt.Println("Too LOW.")
			fmt.Println("-----")
		} else if guess > target {
			fmt.Println("-----")
			fmt.Println("Too HIGH.")
			fmt.Println("-----")
		} else {
			success = true
			fmt.Println("~------------~")
			fmt.Println("| YOU WIN!!! |")
			fmt.Println("~------------~")
			break
		}
	}

	if !success {
		fmt.Println("Oh, sorry.")
		fmt.Println("It was", target)
	}
}
