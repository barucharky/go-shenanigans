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
var errString error = errors.New("you don't know your numbers from your letters")

// GetInt gets an integer from the user
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

// GetFloat gets a float from the user
func GetFloat() (float64, error) {

	// Call GetFloat
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

// GetInput gets input from the user, returns trimmed string
func GetInput() string {

	// Initialize input from keyboard
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	// Get input
	var input string
	var err error

	input, err = reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	// Trim string
	input = strings.TrimSpace(input)

	return input
}

// ConvertInt converts a string to an integer
func ConvertInt(s string) (int, error) {

	// Attempt conversion
	var theInt int
	var err error

	theInt, err = strconv.Atoi(s)

	if err != nil {

		_, err = strconv.ParseFloat(s, 64)

		if err == nil {
			return 0, errFloat
		}

		return 0, errString
	}

	return theInt, nil
}

// ConvertFloat converts a string to a float
func ConvertFloat(s string) (float64, error) {

	// Attempt conversion
	var theFloat float64
	var err error

	theFloat, err = strconv.ParseFloat(s, 64)

	if err != nil {
		return 0, errString
	}

	return theFloat, nil
}
