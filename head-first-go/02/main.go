package main

import (
	"bufio"
	"fmt"
	"math/rand"
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
		fmt.Println("~------------~")
		fmt.Println("I am functioning within established parameters.")
	}

	fmt.Println("~------------~")
	fmt.Println("Reader type:", reflect.TypeOf(reader))
	fmt.Println("Input type:", reflect.TypeOf(input))

	// strings
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println("~------------~")
	fmt.Println("fixed string:")
	fmt.Println(fixed)

	// random number
	// why does this always produce the same number?
	target := rand.Intn(100) + 1
	fmt.Println("~------------~")
	fmt.Println("Random number:")
	fmt.Println(target)

	// loop
	fmt.Println("~------------~")
	for x := 4; x <= 6; x++ {
		fmt.Println("x is now", x)
	}

	fmt.Println("~------------~")
	y := 15
	for y >= 5 {
		fmt.Println("y is now", y)
		y -= 5
	}
}
