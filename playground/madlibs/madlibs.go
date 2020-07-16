// B''H

/*
go mod init sandbox/madlibs
go run madlibs.go
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	// -- ------------------------------
	// Get the sentences and the class
	fmt.Println("Type up until class or end of story, then press enter")

	// Loop until finished, adding to slices
	var sentences []string
	var classes []string

	for true {

		fmt.Println("Type your sentence:")
		var sentence string = getString()

		if sentence == "" {
			break
		}

		fmt.Println("Enter class of word:")
		var class string = getString()

		if class == "" {
			break
		}

		sentences = append(sentences, sentence)
		classes = append(classes, class)

	}

	// -- ------------------------------
	// Get words from second user
	// Clear screen
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	// Loop through classes, appending sentences and user input to story
	var story string

	for i, c := range classes {
		fmt.Printf("%s:\n", c)
		word := getString()

		story = fmt.Sprintf("%s %s %s", story, sentences[i], word)
	}

	// Show the whole story
	fmt.Println("Here's the whole story:\n")
	fmt.Println(story)

}

func getString() string {

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(input)
}

func getClass() string {

	// Define classes
	classes := []string{
		"adjective",
		"noun",
		"adverb",
		"verb",
	}

	// Print options
	for i, s := range classes {
		fmt.Printf("%d: %s\n", i, s)
	}

	var selection int

	// Get user input
	for true {

		// Get input from user
		var input string = getString()

		// If user entered nothing (just pressed return) return empty string
		if input == "" {
			return ""
		}

		// Convert to int
		var err error
		selection, err = strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if selection > len(classes)-1 {
			fmt.Println("Invalid selection")
			continue
		} else {
			break
		}
	}

	return classes[selection]

}
