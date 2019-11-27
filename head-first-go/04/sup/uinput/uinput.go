package uinput

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

var errFloat error = errors.New("that was never an option")
var errString error = errors.New("learn the difference between a letter and a number")

// GetInt ets an string from user and converts to integer
func GetInt() (int, error) {

	// Get input from user
	var doodad string = GetInput()

	// Convert to int
	var theInt int
	var err error

	theInt, err = ConvertInt(doodad)
	if err != nil {
		return 0, err
	}

	return theInt, nil
}

// GetFloat ets an string from user and converts to float
func GetFloat() (float64, error) {

	// Get input from user
	var doodad string = GetInput()

	// Convert to Float
	var theFloat float64
	var err error

	theFloat, err = ConvertFloat(doodad)
	if err != nil {
		return 0, err
	}

	return theFloat, nil
}

// GetInput Gets input and trims whitespaces
func GetInput() string {

	var input string
	var err error
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	return input
}

// ConvertInt Converts a string to an integer
func ConvertInt(s string) (int, error) {

	var number int
	var err error

	number, err = strconv.Atoi(s)
	if err != nil {
		_, err = strconv.ParseFloat(s, 64)
		if err == nil {
			return 0, errFloat
		}

		return 0, errString
	}

	return number, nil
}

// ConvertFloat Converts a string to a float
func ConvertFloat(s string) (float64, error) {

	var number float64
	var err error

	number, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errString
	}

	return number, nil
}
