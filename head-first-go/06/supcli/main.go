package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/english"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/french"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/hebrew"
	"github.com/barucharky/go-shenanigans/head-first-go/04/sup/langs/spanish"
)

var errType error = errors.New("That aint no language, fool")
var errNull error = errors.New("ENGLISH. Do. You. Speak. It. Pick a language.")

func main() {

	var arguments []string = os.Args[1:]

	if len(arguments) == 0 {
		log.Fatal(errNull)
	}

	var argument string
	for _, argument = range arguments {
		greeting(argument)
	}

}

func greeting(argument string) {
	if argument == "en" {
		english.Hello()
		english.How()
	} else if argument == "fr" {
		french.Bonjour()
		french.Comm()
	} else if argument == "he" {
		hebrew.Shalom()
		hebrew.Mah()
	} else if argument == "sp" {
		spanish.Hola()
		spanish.Como()
	} else if argument == "usage" {
		usage()
	} else {
		log.Fatal(errType)
	}
}

func usage() {
	fmt.Println("Use the following abbreviations:")
	fmt.Println("en : English")
	fmt.Println("fr : French")
	fmt.Println("he : Hebrew")
	fmt.Println("sp : Spanish")
}
