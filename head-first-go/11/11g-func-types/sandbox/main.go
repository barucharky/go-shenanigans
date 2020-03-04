// B"H

package main

import "fmt"

// -- --------------------------------------
// You can declare a function signature as a type:
type myFunctionType = func(a string, b string) string

func wazzup(fn myFunctionType) {

	var newStr = fn("hello", "world")

	fmt.Println(newStr)
}

// -- --------------------------------------
func main() {

	// -- ----------------------------------
	// You can be explicit about what func type your func is going to implement:

	var implementation1 myFunctionType = func(a string, b string) string {
		return fmt.Sprintf("%s+%s", a, b)
	}

	wazzup(implementation1)

	// -- ----------------------------------
	// You can be implicit:

	implementation2 := func(a string, b string) string {
		return fmt.Sprintf("%s to the %s", a, b)
	}

	wazzup(implementation2)

	// -- ----------------------------------
	// You can even be anonymous:

	wazzup(
		func(a string, b string) string {
			return fmt.Sprintf("%s, I want to say %s", b, a)
		},
	)
}
