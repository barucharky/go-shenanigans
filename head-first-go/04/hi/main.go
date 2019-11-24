// B"H

package main

import (
	"fmt"
	"log"

	"github.com/barucharky/go-shenanigans/head-first-go/04/greeting"
	"github.com/barucharky/go-shenanigans/head-first-go/04/greeting/dansk"
	"github.com/barucharky/go-shenanigans/head-first-go/04/greeting/deutsch"
	"github.com/barucharky/go-shenanigans/head-first-go/04/keyboard"
)

func main() {

	fmt.Print("Choose a language:\n")
	fmt.Print("1 : English\n")
	fmt.Print("2 : Dansk\n")
	fmt.Print("3 : Deutsch\n")

	// -- -----------------------------
	lang, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}
	// -- -----------------------------

	// -- -----------------------------
	if lang == 1 {
		greeting.Hello()
		greeting.Hi()
	} else if lang == 2 {
		dansk.Hej()
		dansk.GodMorgen()
	} else {
		deutsch.Hallo()
		deutsch.GutenTag()
	}
	// -- -----------------------------

}
