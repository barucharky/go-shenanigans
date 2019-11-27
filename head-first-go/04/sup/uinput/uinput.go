package uinput

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetInt gets user input, gives back an integer
func GetInt() (int, error) {

	var doodad string = getInput()

	var err error
	var theInt int
	theInt, err = convert(doodad)
	if err != nil {
		return 0, err
	}

	return theInt, nil
}

func getInput() string {

	// Initialize input from keyboard
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	// Get input
	var input string
	var err error

	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	return input
}

func convert(s string) (int, error) {

	// Declare errors
	var floatErr error = errors.New("not a real choice, bub. try an integer")
	var stringErr error = errors.New("learn the difference between a letter and a number")

	// Convert to int
	var number int
	var err error

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
