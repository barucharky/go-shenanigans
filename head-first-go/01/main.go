package main

import (
	"fmt"
	"reflect"
)

func main() {
	// -- -----------------
	fmt.Println("---------------------")
	fmt.Println("-   SIMPLE OUTPUT   -")
	fmt.Println("-                   -")
	fmt.Println("word")

	// -- -----------------
	fmt.Println("---------------------")
	fmt.Println("- STRINGS AND RUNES -")
	fmt.Println("-                   -")
	fmt.Println("A   :", 'A')
	fmt.Println("a   :", 'a')
	fmt.Println("漢   :", '漢')

	// -- -----------------
	fmt.Println("---------------------")
	fmt.Println("-       TYPES       -")
	fmt.Println("-                   -")
	fmt.Println("integer       :", 11)
	fmt.Println("floating point:", 11.5)
	fmt.Println("boolean       :", false)
	fmt.Println("string        :", "This is a string")

	fmt.Println("------------------")
	fmt.Println("The type for 11                 :", reflect.TypeOf(11))
	fmt.Println("The type for 11.5               :", reflect.TypeOf(11.5))
	fmt.Println("The type for false              :", reflect.TypeOf(false))
	fmt.Println("The type for `This is a string` :", reflect.TypeOf("This is a string"))
	fmt.Println("The type for '漢'               :", reflect.TypeOf('漢'))

	// -- ---------------------
	fmt.Println("------------------------")
	fmt.Println("- VARIABLES APPROACH 1 -")
	fmt.Println(" (declare, assign, use) ")
	fmt.Println("-                      -")

	// declare:
	var firstName string
	// assign:
	firstName = "Moshe"
	// use:
	fmt.Println(firstName)

	// declare multiple:
	var houseCost, carCost float64
	// assign multiple:
	houseCost, carCost = 1789.45, 894.89
	// use:
	fmt.Println(houseCost, carCost)

	// -- ---------------------
	fmt.Println("--------------------------")
	fmt.Println("-  VARIABLES APPROACH 2  -")
	fmt.Println(" (declare & assign, use)  ")
	fmt.Println("-                        -")

	// Declare and assign:
	var petName string = "Rex"
	// Use:
	fmt.Println(petName)

	// Declare and assign multiple:
	var isExpensive, isBlue bool = true, false
	// use:
	fmt.Println(isExpensive, isBlue)

	// -- ---------------------
	fmt.Println("--------------------------")
	fmt.Println("-  VARIABLES APPROACH 2b -")
	fmt.Println(" (declare & assign, use)  ")
	fmt.Println("-without mentioning type -")
	fmt.Println("-                        -")
	fmt.Println("-                        -")
	fmt.Println("-    can be dangerous    -")
	fmt.Println("-                        -")

	// usage

	var boatName = "O'Neal"
	fmt.Println(boatName)

	// -- ---------------------
	fmt.Println("--------------------------")
	fmt.Println("-  VARIABLES APPROACH 3  -")
	fmt.Println("     (assign, use)        ")
	fmt.Println("-                        -")
	fmt.Println("-        called:         -")
	fmt.Println("-short variable declaration-")
	fmt.Println("-                        -")

	// assign
	lastName := "Shmorgisborg"
	//use
	fmt.Println(lastName)
}
