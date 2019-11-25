// Package uinput is for reading and converting to a float
package uinput

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat reads a float from the user
// It returns the number or an error
func GetFloat() (float64, error) {

	// Initialize
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	var input string
	var err error

	// Get input
	input, err = reader.ReadString('\n')

	input = strings.TrimSpace(input)

	// Convert to float
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}
