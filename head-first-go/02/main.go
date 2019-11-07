package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	// input
	reader := bufio.NewReader(os.Stdin)
	question := "how are you\n"

	fmt.Println("type something")
	input, _ := reader.ReadString('\n')

	if input == question {
		fmt.Println("\nI am functioning within established parameters.\n")
	}

	fmt.Println("Reader type:", reflect.TypeOf(reader))
	fmt.Println("Input type:", reflect.TypeOf(input))

	// strings
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println("\nfixed string:")
	fmt.Println(fixed)
}
