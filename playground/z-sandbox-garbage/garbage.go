// B''H

/*
go mod init sandbox/garbage
go run garbage.go
*/

package main

import (
	"fmt"
)

func helper(s string, err error) string {

	if err != nil {
		panic(err)
	}

	return s
}

func errFunc(b bool) (string, error) {

	if b == false {
		return "", fmt.Errorf("this is the error")
	}

	return "this is the string", nil
}

func main() {

	var message string

	message = helper(errFunc(true))
	fmt.Println(message)

	message = helper(errFunc(false))
	fmt.Println(message)
}
