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

	// Pick a random number
	var seconds int64 = time.Now().Unix()
	rand.Seed(seconds)
	var target int = rand.Intn(100) + 1

	// Initialize input from keyboard
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	// Start game
	var success bool = false
	var guesses int64

	// Start the loop
	for guesses = 0; guesses < 10; guesses++ {

		fmt.Println("You have", 10-guesses, "left.")
		fmt.Print("Enter your guess: ")

		// Get input from user
		var input string
		var err error

		input, err = reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		// Convert input to int
		var guess int

		guess, err = strconv.Atoi(input)

		if err != nil {
			guesses--
			printMessage("Invalid guess")
			continue
		}

		// Check answer

		if guess < target {
			printMessage("Too LOW")
		} else if guess > target {
			printMessage("Too HIGH")
		} else {
			success = true
			printMessage("YOU WINNNN!!!")
			break
		}
	}

	if !success {
		fmt.Println("You lose.")
		fmt.Println("It was", target)
	}

}

func printMessage(message string) {
	fmt.Println("-----")
	fmt.Println(message)
	fmt.Println("-----")
}
