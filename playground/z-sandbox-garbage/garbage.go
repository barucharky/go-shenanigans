// B''H

/*
go mod init sandbox/garbage
go run garbage.go
*/

package main

import (
	"fmt"
)

func main() {

	fmt.Println("testing return to break loop")

	myStr := myLoop()
	fmt.Println(myStr)
}

func myLoop() string {

	for i := 0; i < 10; i++ {
		if i == 11 {
			return "thing"
		}
	}

	return "unbroken"
}
