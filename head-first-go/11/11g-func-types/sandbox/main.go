// B"H

package main

import "fmt"

// Declare the function type's signature
type myFunctionType = func(a string, b string) string

// Write a function to implement the functions
func wazzup(fn myFunctionType) {

	var newStr string = fn("hello", "world")

	fmt.Println(newStr)
}

// Here's an implementation to use later
func implementation3(a string, b string) string {

	return "ashemoooooorrgissBBOOOOOOORDddah!!!!"
}

func main() {

	// Explicit implementation
	var implementation1 myFunctionType = func(a string, b string) string {

		return fmt.Sprintf("%s %s", a, b)
	}

	fmt.Println("Implementation 1:")

	wazzup(implementation1)

	fmt.Println()

	// -- ---------------------------
	// Implicit
	implementation2 := func(a string, b string) string {
		return fmt.Sprintf("%s %s", b, a)
	}

	fmt.Println("Implementation 2:")

	wazzup(implementation2)

	fmt.Println()

	// -- ---------------------------
	// More implicit
	fmt.Println("Implementation 3:")

	wazzup(implementation3)

	fmt.Println()

	// -- ---------------------------
	// Anonymous
	fmt.Println("Anonymous:")

	wazzup(
		func(a string, b string) string {
			return fmt.Sprintf("%s %s!!!!", a, b)
		},
	)

}
