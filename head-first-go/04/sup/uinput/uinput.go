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

	// Call GetInput to get a string
	var doodad string = GetInput()

	// call ConvertInt to convert string to int
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

	// Call GetInput to get a string
	var doodad string = GetInput()

	// Call ConvertFloat to conver string to float
	var theFloat float64
	var err error

	theFloat, err = ConvertFloat(doodad)

	if err != nil {
		return 0, err
	}

	return theFloat, nil
}

// GetInput gets input from the user
func GetInput() string {

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

// ConvertInt converts a string to an integer
func ConvertInt(s string) (int, error) {

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

	var theFloat float64
	var err error

	theFloat, err = strconv.ParseFloat(s, 64)

	if err != nil {
		return 0, errString
	}

	return theFloat, nil
}
