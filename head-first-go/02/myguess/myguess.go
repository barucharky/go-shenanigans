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

	// -- ----------------------
	// produce random number
	var seconds int64 = time.Now().Unix()
	rand.Seed(seconds)
	var target int = rand.Intn(100) + 1

	fmt.Println("I've picked a number between 1 and 100")
	fmt.Println("See if you can guess what it is.")
	// -- ----------------------

	// -- ----------------------
	// initiate input from keyboard
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	// -- ----------------------

	// -- ----------------------
	// Start game

	var success bool = false

	// Start loop
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "left.")
		fmt.Print("Enter your guess: ")

		// get user input
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// trim whitespace
		input = strings.TrimSpace(input)

		// attempt conversion to int
		guess, err := strconv.Atoi(input)

		// if fail, give another turn
		if err != nil {
			guesses--
			fmt.Println("Oops. Try again with a number")
			continue
		}

		// check answer
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
			fmt.Println("~-----------------------~")
			fmt.Println("| That's it! You won!!! |")
			fmt.Println("~-----------------------~")
			break
		}
	}

	if !success {
		fmt.Println("Ooooh, too bad.")
		fmt.Println("It was", target)
	}
}
