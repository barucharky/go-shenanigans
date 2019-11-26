package uinput

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (int, error) {

	// Get input from user
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	var input string
	var err error
	input, err = reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	// Trim spaces
	input = strings.TrimSpace(input)

	// Convert to int
	var number int
	number, err = strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return number, nil
}
