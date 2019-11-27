package main

import (
	"fmt"

	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/english"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/french"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/hebrew"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/spanish"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/uinput"
)

func main() {

	// Show options
	fmt.Println("Choose a language:")
	fmt.Println("1 : English")
	fmt.Println("2 : French")
	fmt.Println("3 : Hebrew")
	fmt.Println("4 : Spanish")

	// Get user's selection

	var selection int
	var err error

	for true {
		selection, err = uinput.GetInt()
		if err == nil {
			break
		}

		println(err.Error())
	}

	greeting(selection)
}

func greeting(number int) {
	if number == 1 {
		english.Hello()
		english.How()
	} else if number == 2 {
		french.Bonjour()
		french.Comm()
	} else if number == 3 {
		hebrew.Shalom()
		hebrew.Mah()
	} else if number == 4 {
		spanish.Hola()
		spanish.Como()
	} else {
		fmt.Println("Can't say I know that one.")
	}
}
