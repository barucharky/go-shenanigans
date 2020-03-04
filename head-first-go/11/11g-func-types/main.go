// B"H

package main

import "fmt"

// -- --------------------------------------
// You can declare a function signature as a type:
type myFunctionType = func(a string, b string) string

// -- --------------------------------------
// Notice how this function takes in a function type as a parameter:
func wazzup(fn myFunctionType) {

	var newStr string = fn("hello", "world")

	fmt.Println(newStr)
}

// -- --------------------------------------
// Notice, how the function signature matches myFunctionType:
func implementation3(a string, b string) string {

	return "What the shmorgisborg!?"

}

// -- --------------------------------------
func main() {

	// -- ----------------------------------
	//             APPROACH 1
	// You can be explicit about what func type your func is going to implement:
	var implementation1 myFunctionType = func(a string, b string) string {
		return fmt.Sprintf("%s %s", a, b)
	}

	wazzup(implementation1)

	// -- ----------------------------------
	//             APPROACH 2
	// You can be a tad more implicit:
	implementation2 := func(a, b string) string {
		return fmt.Sprintf("%s %s", b, a)
	}

	wazzup(implementation2)

	// -- ----------------------------------
	//             APPROACH 3
	// In general, I like this approach the best.
	//
	// You can be even more implicit:
	wazzup(implementation3)

	// -- ----------------------------------
	//             APPROACH 4
	// You can even be anonymous:
	wazzup(

		func(a, b string) string {
			return fmt.Sprintf("%s %s!", a, b)
		},
	)

}
