package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var errUsage error = errors.New("invalid argument. see usage")

func avg(numbers ...float64) float64 {

	var sum float64

	sum = add(numbers...)

	return sum / float64(len(numbers))
}

func add(numbers ...float64) float64 {

	var sum float64
	var number float64

	for _, number = range numbers {
		sum += number
	}

	return sum
}

func main() {

	var mathType string = os.Args[1]
	var arguments []string = os.Args[2:]
	var numbers []float64

	var argument string
	for _, argument = range arguments {

		var number float64
		var err error

		number, err = strconv.ParseFloat(argument, 64)

		if err != nil {
			log.Fatal(errUsage)
		}

		numbers = append(numbers, number)
	}

	if mathType == "usage" {
		usage()
	} else if mathType == "add" {
		fmt.Printf("The sum of your numbers is %0.2f\n", add(numbers...))
	} else if mathType == "avg" {
		fmt.Printf("The average of your numbers is %0.2f\n", avg(numbers...))
	}
}

func usage() {
	fmt.Println("Now here's what you're gonna do.")
	fmt.Println("You're gonna type like this:")
	fmt.Println("mymath [add|avg|usage] [arguments]")
	fmt.Println("That'll give you what you want")
}
