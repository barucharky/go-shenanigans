// B"H

package main

import "fmt"

// -- -----------------------------------
// Declare the function type's signature
type myFunctionType = func(a, b string) string

// -- -----------------------------------
// Write a function to implement the functions
func wazzap(fn myFunctionType) {
	newStr := fn("hello", "world")

	fmt.Println(newStr)
}

// -- -----------------------------------
// Here's an implementation to use later
func implementation3(a, b string) string {
	return "a-shemooorgisBOOOOOOrda!!!"
}

// -- -----------------------------------
func main() {

	// -- -------------------------------
	// Explicit implementation
	fmt.Println("Explicit:")

	var implementation1 myFunctionType = func(a, b string) string {
		return fmt.Sprintf("%s %s", a, b)
	}

	wazzap(implementation1)

	fmt.Println()

	// -- -------------------------------
	// Implicit
	fmt.Println("Implicit:")

	implementation2 := func(a, b string) string {
		return fmt.Sprintf("%s %s", b, a)
	}

	wazzap(implementation2)

	fmt.Println()

	// -- -------------------------------
	// More implicit
	fmt.Println("More implicit")

	wazzap(implementation3)

	fmt.Println()

	// -- -------------------------------
	// Anonymous
	fmt.Println("You can even be anonymous:")

	wazzap(
		func(a string, b string) string {
			return fmt.Sprintf("%s %s!!!", a, b)
		},
	)

}
