package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Get a random number
	var target int
	target = pickNumber()

	// Start game
	var success bool = false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "left.")
		fmt.Print("Enter your guess: ")

		// Get input from user
		var guess int
		var err error
		guess, err = getGuess()
		if err != nil {
			guesses--
			printMessage("Must be an integer.")
			continue
		}

		// Check answer
		var response string
		response, success = checkGuess(guess, target)
		printMessage(response)
		if success {
			break
		}
	}

	if !success {
		printMessage("You lost.")
		fmt.Println("It was", target)
	}
}

func pickNumber() int {
	var seconds int64 = time.Now().Unix()
	rand.Seed(seconds)
	var number int = rand.Intn(100) + 1

	fmt.Println("I've chosen a number between 1 and 100.")
	fmt.Println("Can you guess what it is?")
	fmt.Println("-----")
	return number
}

func getGuess() (int, error) {

	// Initialize input from keyboard
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	// Get user input
	var input string
	var err error
	input, err = reader.ReadString('\n')

	// Trim whitespaces
	input = strings.TrimSpace(input)

	// Attempt conversion
	var guess int
	guess, err = strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("no")
	}

	return guess, nil

}

func printMessage(message string) {
	fmt.Println("-----")
	fmt.Println(message)
	fmt.Println("-----")
}

func checkGuess(guess, target int) (string, bool) {
	if guess < target {
		return "Too LOW", false
	} else if guess > target {
		return "Too HIGH", false
	}

	return "You Win!", true
}
