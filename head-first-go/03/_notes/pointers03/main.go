// B''H

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// -- ------------------------------
	var amount1 int64 = 6
	var amount2 int64 = amount1
	var amount3pointer *int64 = &amount1

	// -- ------------------------------
	fmt.Println("----------")
	fmt.Println("amount1")
	fmt.Println(amount1)
	fmt.Println(&amount1)

	fmt.Println(reflect.TypeOf(amount1))
	fmt.Println(reflect.TypeOf(&amount1))

	// -- -------------------------------
	fmt.Println("----------")
	fmt.Println("amount2")
	fmt.Println(amount2)
	fmt.Println(&amount2)

	fmt.Println(reflect.TypeOf(amount2))
	fmt.Println(reflect.TypeOf(&amount2))

	// -- -------------------------------
	fmt.Println("----------")
	fmt.Println("amount3pointer")
	fmt.Println(amount3pointer)

	fmt.Println(reflect.TypeOf(amount3pointer))
	fmt.Println(reflect.TypeOf(*amount3pointer))

	// -- -------------------------------
	fmt.Println("----------")
	fmt.Println("Change the value of the pointer")
	*amount3pointer = 9999
	fmt.Println(*amount3pointer)
	fmt.Println(amount1)
	fmt.Println(&amount3pointer)

}
