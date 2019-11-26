package uinput

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (int, error) {

	// Get input from user
	var input string
	var err error
	input, err = getInput()
	if err != nil {
		return 0, err
	}

	// Convert to int
	var number int

	number, err = convert(input)
	if err != nil {
		return 0, err
	}

	return number, nil
}

// getInput gets a string from the user
func getInput() (string, error) {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var input string
	var err error

	// Get the input
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Trim spaces
	input = strings.TrimSpace(input)

	return input, nil
}

// convert attempts to convert a string to an int
func convert(s string) (int, error) {
	var number int
	var err error
	
	var floatErr error = errors.New("You entered a float. Only an integer is acceptable.")
	var stringErr error = errors.New("You entered a string. Only an integer is acceptable.")
	
	// Attempt conversion to int
	number, err = strconv.Atoi(s)

	if err != nil {
		_, err = strconv.ParseFloat(s, 64)
		if err == nil {
			return 0, floatErr
			} else {
				return 0, stringErr
		}
	}

	return number, nil
}
