// B"H

package main

import "fmt"

// Declare the function type's signature
type myFunctionType = func(a string, b string) string

// Write a function to implement the functions
func wazzup(fn myFunctionType) {

	var newString string = fn("hello", "world")

	fmt.Println(newString)
}

// Here's an implementation to use later
func implementation3(a string, b string) string {
	return "ashemoooooorrgissBBOOOOOOORDddah!!!!"
}

func main() {

	// Explicit implementation
	fmt.Println("Explicit:")

	var implementation1 myFunctionType = func(a string, b string) string {
		return fmt.Sprintf("%s %s", a, b)
	}

	wazzup(implementation1)

	fmt.Println()

	// -- ---------------------------
	// Implicit
	fmt.Println("Implicit:")

	implementation2 := func(a string, b string) string {
		return fmt.Sprintf("%s %s", b, a)
	}

	wazzup(implementation2)

	fmt.Println()

	// -- ---------------------------
	// More implicit
	fmt.Println("More implicit:")

	wazzup(implementation3)

	fmt.Println()

	// -- ---------------------------
	// Anonymous
	fmt.Println("Anonymous:")

	wazzup(
		func(a string, b string) string {
			return fmt.Sprintf("%s %s!", a, b)
		},
	)

}
