// B''H

/*
go mod init sandbox/sandbox
go run main.go
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var search *string = flag.String("s", "", "If you want to search for a file or dir, enter the name")
var dir *string = flag.String("d", "", "Specify directory to search")
var list *bool = flag.Bool("l", false, "List contents before search")

func fileSearch(f []os.FileInfo) {

	var foundMatch bool

	for _, file := range f {
		if strings.ToLower(*search) == strings.ToLower(file.Name()) {
			foundMatch = true
		}
	}

	if foundMatch {
		fmt.Println("Well, I'll be... Found Match!")
	} else {
		fmt.Println("Sorry, mate. No match found.")
	}
}

func getFiles() ([]os.FileInfo, string, error) {

	// Declare relevant variables
	var directory string
	var err error

	if *dir == "" {
		directory, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		directory = *dir
	}

	// Read all the file names in the directory
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, "", err
	}

	return files, directory, nil
}

func main() {

	flag.Parse()

	// Get a list of the specified directory's contents

	files, directory, err := getFiles()
	if err != nil {
		log.Fatal(err)
	}

	// Print directory contents if specified

	if *list {
		fmt.Printf("Directory %s contains:\n", directory)
		for _, f := range files {
			fmt.Println(f.Name())
		}

		fmt.Println("-----")
	}

	// Run search if requested
	if *search != "" {
		fmt.Println("Running search...")
		fileSearch(files)
	} else {
		fmt.Println("Search parameter needed.")
	}
}
