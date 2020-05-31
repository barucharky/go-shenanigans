package main

import "fmt"

func addOne(ogInt *int) *int {
	fmt.Println(ogInt)
	*ogInt = *ogInt + 1
	return ogInt
}

func main() {
	var ourNum int
	var ourNumP *int
	ourNumP = addOne(&ourNum)

	fmt.Println(ourNumP)
	fmt.Println(&ourNum)
}
