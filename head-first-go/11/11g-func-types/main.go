// B"H

package main

import "fmt"

// -- --------------------------------------
// You can declare a function signature as a type:
type myFunctionType = func(a string, b string) string

func wazzup(fn myFunctionType) {

	var newStr string = fn("hello", "world")

	fmt.Println(newStr)
}

// -- --------------------------------------
func main() {

	// -- ----------------------------------
	// You can be explicit about what func type your func is going to implement:
	var implementation1 myFunctionType = func(a string, b string) string {
		return fmt.Sprintf("%s %s", a, b)
	}

	wazzup(implementation1)

	// -- ----------------------------------
	// You can be implicit:
	implementation2 := func(a, b string) string {
		return fmt.Sprintf("%s %s", b, a)
	}

	wazzup(implementation2)

	// -- ----------------------------------
	// You can even be anonymous:
	wazzup(

		func(a, b string) string {
			return fmt.Sprintf("%s %s!", a, b)
		},
	)

}
