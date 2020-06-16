package main

import "fmt"

// B''H

/*
go mod init sandbox/garbage
go run garbage.go
*/

func main() {
	var listA []string = []string{"a", "b", "c"}
	fmt.Println(listA)

	for i, c := range listA {
		fmt.Println("Index:", i)
		fmt.Println("Content", c)
	}
}
