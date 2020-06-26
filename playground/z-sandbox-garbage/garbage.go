// B''H

/*
go mod init sandbox/garbage
go run garbage.go
*/

package main

import (
	"fmt"
)

type myStruc struct {
	number int
	name   string
	isMe   bool
}

func retPointer() *myStruc {
	var theStruct myStruc = myStruc{number: 1, name: "Baruch", isMe: true}

	return &theStruct
}

func main() {

	newStruct := retPointer()

	fmt.Println(newStruct)
}
