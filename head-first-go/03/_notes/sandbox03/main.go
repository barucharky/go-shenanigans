package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// -- -------------------------------------
func double(num float64) float64 {
	return num * 2
}

// -- -------------------------------------
func sayHello(name string, saying string) (output string, err error) {
	if name == "baruch" {
		return saying, err
	} else {
		return "", fmt.Errorf("You're not Baruch, %s", name)
	}
}

func main() {

	// -- ---------------------------------

	// Get a stdin
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// -- ---------------------------------
	// trim and convert to float64
	input = strings.TrimSpace(input)

	inputNum, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	// -- ---------------------------------
	fmt.Println(double(inputNum))

	// -- ---------------------------------
	// Utilizing custom error
	var myErr error = errors.New("No way, no how")
	fmt.Println(myErr)
	// log.Fatal(myErr)

	// -- ---------------------------------
	// Utilizing custom error with formatting
	var myErr2 error = fmt.Errorf("No way, no how %s", "Jack")
	fmt.Println(myErr2)
	// log.Fatal(myErr2)

	fmt.Println("-----")

	// -- ---------------------------------
	// Get output without error
	myOutput, err := sayHello("baruch", "hellooooo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(myOutput)

	fmt.Println("-----")

	// -- ---------------------------------
	// Get output with error
	myOutput, err = sayHello("mendy", "hellooooo")
	if err != nil {
		log.Fatal(err)
	}
}
