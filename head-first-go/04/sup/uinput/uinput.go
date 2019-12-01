package uinput

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

// Declare errors
var errFloat error = errors.New("that was never an option")
var errString error = errors.New("seems you don't know your letters from your numbers")

// GetInt gets input from user and converts to int
func GetInt() (int, error) {

	// Call GetInput
	var doodad string

	doodad = GetInput()

	// Call ConvertInt
	var theInt int
	var err error

	theInt, err = ConvertInt(doodad)
	if err != nil {
		return 0, err
	}

	return theInt, nil
}

// GetFloat gets input from user and converts to float
func GetFloat() (float64, error) {

	// Call GetInput
	var doodad string

	doodad = GetInput()

	// Call ConvertFloat
	var theFloat float64
	var err error

	theFloat, err = ConvertFloat(doodad)
	if err != nil {
		return 0, err
	}

	return theFloat, nil
}

// GetInput gets input from user and trims it
func GetInput() string {

	// Initialize input from keyboard
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	var input string
	var err error

	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	return input
}

// ConvertInt converts a sting to an int
func ConvertInt(s string) (int, error) {

	// Attempt conversion
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

// ConvertFloat converts a sting to a float
func ConvertFloat(s string) (float64, error) {

	// Attempt conversion
	var number float64
	var err error

	number, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errString
	}

	return number, nil
}
