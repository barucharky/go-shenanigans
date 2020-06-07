// B''H

/*
go mod init sandbox/garbage
go run garbage.go
*/

package main

import "fmt"

func main() {
	s := "hello, world"
	hello := s[:5]
	world := s[7:]

	fmt.Println(s)
	fmt.Println(hello)
	fmt.Println(world)

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for _, v := range s {
		fmt.Printf("%q", v)
	}

	for i, r := range "Hello, בה ᒂ 🌼" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

}
