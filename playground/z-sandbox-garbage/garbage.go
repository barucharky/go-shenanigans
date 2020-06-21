// B''H

/*
go mod init sandbox/garbage
go run garbage.go
*/
package main

import (
	"fmt"
)

/*
Just a little review of how functions work. You'll recognize a function because it's a command with parenthesis after it
*/

// A function can receive values and return values
func receiveReturn(valueReceived string) string {

	// The variable valueReceived is a copy of the value passed in

	fmt.Println("This is the value I received:")
	fmt.Println(valueReceived)

	return "Value Returned"

}

// A function doesn't have to receive or return. It can also do one or the other
func noReceiveNoReturn() {
	fmt.Println("I received nothing and returned nothing")
}

func main() {
	var valuePassedIn string = "Here I am"
	var valueReturned string

	fmt.Println("Here are the values before they are passed in and received:")
	fmt.Println("The value of the variable to be passed in:", valuePassedIn)
	fmt.Println("The value of the variable that will be returned:", valueReturned)

	// -- ----------------------------------------------------------
	fmt.Println("---------")

	// The function that received a passed-in value works like this
	valueReturned = receiveReturn(valuePassedIn)

	fmt.Println("\nHere is the value the function returned:")
	fmt.Println(valueReturned)

	// -- -----------------------------------------------------------
	fmt.Println("---------")

	// The function that neither receives nor returns works like this
	noReceiveNoReturn()

}
