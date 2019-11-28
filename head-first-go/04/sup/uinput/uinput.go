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
var errString error = errors.New("seems you don't know letters from numbers")

// GetInt gets an int from the user
func GetInt() (int, error) {

	// Get input from user
	var doodad string = GetInput()

	var theInt int
	var err error

	// Convert to integer
	theInt, err = ConvertInt(doodad)
	if err != nil {
		return 0, err
	}

	return theInt, nil
}

// GetFloat gets a float from the user
func GetFloat() (float64, error) {

	// Get input from user
	var doodad string = GetInput()

	var theFloat float64
	var err error

	// Convert to float
	theFloat, err = ConvertFloat(doodad)
	if err != nil {
		return 0, err
	}

	return theFloat, nil
}

// GetInput gets input from the user
func GetInput() string {

	// Declare variables
	var input string
	var err error
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	// Get input
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Trim whitespaces
	input = strings.TrimSpace(input)
	return input
}

// ConvertInt converts to int
func ConvertInt(s string) (int, error) {

	// Declare variables
	var number int
	var err error

	// Attempt conversion to int
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

// ConvertFloat converts to float
func ConvertFloat(s string) (float64, error) {

	// Declare variables
	var number float64
	var err error

	// Attempt conversion to float
	number, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errString
	}

	return number, nil
}
