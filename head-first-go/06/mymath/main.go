package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var errInvalid error = errors.New("invalid arguments. see usage")

func add(numbers ...float64) float64 {

	var sum float64
	var number float64

	for _, number = range numbers {

		sum += number

	}

	return sum

}

func avg(numbers ...float64) float64 {

	var sum float64 = add(numbers...)

	return sum / float64(len(numbers))
}

func main() {

	// Test number of arguments
	if len(os.Args) < 2 {
		log.Fatal(errInvalid)
	}

	// Get type of math and list of arguments
	var mathType string = os.Args[1]
	var arguments []string = os.Args[2:]

	// Convert arguments to float
	var numbers []float64 = convertArgs(arguments...)

	// Get results
	getResults(mathType, numbers...)

}

func convertArgs(arguments ...string) []float64 {

	var argument string
	var numbers []float64

	for _, argument = range arguments {

		var number float64
		var err error

		number, err = strconv.ParseFloat(argument, 64)

		if err != nil {
			log.Fatal(errInvalid)
		}

		numbers = append(numbers, number)

	}

	return numbers

}

func getResults(mathType string, numbers ...float64) {

	if mathType == "add" || mathType == "sum" {
		fmt.Printf("The sum of your numbers is %0.2f\n", add(numbers...))
	} else if mathType == "avg" || mathType == "average" {
		fmt.Printf("The average of your numbers is %0.2f\n", avg(numbers...))
	} else if mathType == "usage" {
		usage()
	} else {
		log.Fatal(errInvalid)
	}

}

func usage() {
	fmt.Println("Now here's what you're gonna do.")
	fmt.Println("You're gonna type like this:")
	fmt.Println("mymath [add|avg|usage] [arguments]")
	fmt.Println("That'll give you what you want")
}
