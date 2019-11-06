package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Reader type:", reflect.TypeOf(reader))
	fmt.Println("Input type:", reflect.TypeOf(input))
}
