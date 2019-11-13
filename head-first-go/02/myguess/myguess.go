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
	// -- -----------------------------
	// pick a random number
	var seconds int64 = time.Now().Unix()
	rand.Seed(seconds)
	var target int = rand.Intn(100) + 1
	fmt.Println("I've chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	// initialize keyboard input
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	// Play the game
	var success bool = false

	// start the loop
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("-----")
		fmt.Println("You have", 10-guesses, "left.")
		fmt.Print("Enter your guess: ")

		// Get user input
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// trim string
		input = strings.TrimSpace(input)

		// convert to int
		guess, err := strconv.Atoi(input)
		if err != nil {
			guesses--
			fmt.Println("Whoa, that's not a number")
			fmt.Println("Go ahead and try again")
			continue
		}

		// Check answer
		if guess < target {
			fmt.Println("Too LOW")
		} else if guess > target {
			fmt.Println("Too HIGH")
		} else {
			success = true
			fmt.Println("You got it!!!")
			break
		}
	}

	if !success {
		fmt.Println("You lose")
		fmt.Println("It was", target)
	}
}
