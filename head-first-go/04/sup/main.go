package main

import (
	"fmt"
	"log"

	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/english"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/french"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/hebrew"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/spanish"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/uinput"
)

func main() {

	// -- -------------------
	// Menu of languages
	fmt.Println("Choose a language:")
	fmt.Println("1 : English")
	fmt.Println("2 : French")
	fmt.Println("3 : Hebrew")
	fmt.Println("4 : Spanish")
	// -- -------------------

	// -- -------------------
	// Get user input
	var lang float64
	var err error

	lang, err = uinput.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	// Return appropriate greeting
	if lang == 1 {
		english.Hello()
		english.How()
	} else if lang == 2 {
		french.Bonjour()
		french.Comm()
	} else if lang == 3 {
		hebrew.Shalom()
		hebrew.Mah()
	} else if lang == 4 {
		spanish.Hola()
		spanish.Como()
	} else {
		langs.Hey()
		langs.What()
	}
}
