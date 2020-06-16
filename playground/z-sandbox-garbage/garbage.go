// B''H

/*
go mod init sandbox/garbage
go run garbage.go
*/

package main

import (
	"fmt"
)

type employee struct {
	firstName string
	lastName  string
	id        int
	smoker    bool
}

func main() {

	// Assign variable
	empFour := employee{"Monty", "Burns", 234, true}

	var empFourPtr *employee = &empFour

	fmt.Println(empFour)
	fmt.Println(empFourPtr.firstName)
	fmt.Println(empFourPtr)

}
