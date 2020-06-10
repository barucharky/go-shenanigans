// B''H

/*
go mod init sandbox/garbage
go run garbage.go
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func mult(numberList []int) int {
	var answer int = 1
	for _, number := range numberList {
		answer = answer * number
	}
	return answer
}

func main() {
	var numbers []int
	var strings []string

	// Let's put the arguments from the command line into a variable
	strings = os.Args[1:]
	// I called it strings because they are strings at this point.

	// First we have to convert them all into integers

	for _, argument := range strings {
		number, err := strconv.Atoi(argument)

		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, number)
	}

	fmt.Println(mult(numbers))
}
