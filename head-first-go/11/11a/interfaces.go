// B"H

package main

import (
	"fmt"

	"github.com/barucharky/go-shenanigans/head-first-go/11/11a/mypkg"
)

// -- ---------------------------------------
func main() {
	var value mypkg.MyInterface

	value = mypkg.MyType(5)

	value.MethodWithoutParameters()

	value.MethodWithParameter(127.3)

	fmt.Println(value.MethodWithReturnValue())
}
