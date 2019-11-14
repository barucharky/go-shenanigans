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

	// -- --------------------
	// get a random number
	var seconds int64 = time.Now().Unix()
	rand.Seed(seconds)
	var target int = rand.Intn(100) + 1

	fmt.Println("I've picked a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	// initialize reader for keyboard input
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	// -- ---------------------
	// play the game
	var success bool = false

	// start the loop
	for guesses := 0; guesses < 10; guesses++ {

		fmt.Println("You have", 10-guesses, "left.")
		fmt.Print("Enter your guess: ")

		// get user input
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// trim white spaces
		input = strings.TrimSpace(input)

		// convert to int
		guess, err := strconv.Atoi(input)

		// give another turn if conversion fails
		if err != nil {
			guesses--
			fmt.Println("Must be a number. Try again.")
			continue
		}

		// check answer
		if guess < target {
			fmt.Println("---Too LOW---")
		} else if guess > target {
			fmt.Println("---Too HIGH---")
		} else {
			success = true
			fmt.Println("~------------~")
			fmt.Println("| YOU WON!!! |")
			fmt.Println("~------------~")
			break
		}
	}

	if !success {
		fmt.Println("You lose.")
		fmt.Println("It was", target)
	}
}
