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
	menu()

	// Get integer from user
	var selection int
	var err error

	for true {
		selection, err = uinput.GetInt()
		if err == nil {
			break
		}

		fmt.Println("-----")
		fmt.Println(err)
		fmt.Println("-----")
		menu()
	}
	// Show greeting
	greeting(selection)
}

func menu() {
	fmt.Println("Choose a language:")
	fmt.Println("1 : English")
	fmt.Println("2 : French")
	fmt.Println("3 : Hebrew")
	fmt.Println("4 : Spanish")
}

func greeting(selection int) {
	if selection == 1 {
		english.Hello()
		english.How()
	} else if selection == 2 {
		french.Bonjour()
		french.Comm()
	} else if selection == 3 {
		hebrew.Shalom()
		hebrew.Mah()
	} else if selection == 4 {
		spanish.Hola()
		spanish.Como()
	} else {
		fmt.Println("Sorry, I don't know that one")
	}
}
