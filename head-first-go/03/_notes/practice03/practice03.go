package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// set a beginning value
	var name string = "beginning"
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Println("The value of name is", name)
	fmt.Print("Enter a new value: ")

	newOne, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// change the value
	changeValue(&name, newOne)

	fmt.Println("The value of `name` is now", name)
}

func changeValue(change *string, newVal string) {
	*change = newVal
}
