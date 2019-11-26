package main

import (
	"fmt"
	"log"

	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/english"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/french"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/hebrew"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/spanish"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/uinput"
)

func main() {

	fmt.Println("Choose a language for your greeting.")
	fmt.Println("1 : English")
	fmt.Println("2 : French")
	fmt.Println("3 : Hebrew")
	fmt.Println("4 : Spanish")

	var choice int
	var err error

	choice, err = uinput.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	// Show appropriate greeting
	if choice == 1 {
		english.Hello()
		english.How()
	} else if choice == 2 {
		french.Bonjour()
		french.Comm()
	} else if choice == 3 {
		hebrew.Shalom()
		hebrew.Mah()
	} else if choice == 4 {
		spanish.Hola()
		spanish.Como()
	} else {
		fmt.Println("Hmm...I don't know that one.")
	}
}
