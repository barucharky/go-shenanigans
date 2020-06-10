// B''H

/*
go mod init sandbox/sandbox
go run main.go
*/

package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var results []string

var nonRecurs *bool = flag.Bool("n", false, "Non-recursive search")
var fileOnly *bool = flag.Bool("f", false, "Match files only")
var dirOnly *bool = flag.Bool("d", false, "Match directories only")
var dir *string = flag.String("D", "", "Specify search directory")

var errParam error = errors.New("you didn't tell me what to search for")

func main() {
	flag.Parse()

	// Make sure there is a search string
	var search string

	if flag.Args() == nil {
		log.Fatal(errParam)
	} else {
		search = strings.Join(flag.Args(), " ")
	}

	// -- ---------------------------------
	// Establish the working directory
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

	// -- ---------------------------------
	// Run the search

	runSearch(directory, search)

	// Print results
	fmt.Println(results)

}

func runSearch(directory, search string) {

	// Get the list of files in the search directory
	files := fetchFiles(directory)
	var dirBool bool

	for _, file := range files {
		dirBool = dirTest(directory + file.Name())
		if strings.ToLower(search) == strings.ToLower(file.Name()) {
			if *fileOnly && !dirBool {
				results = append(results, directory+file.Name())
			} else if *dirOnly && dirBool {
				results = append(results, directory+file.Name())
			} else if !*fileOnly && !*dirOnly {
				results = append(results, directory+file.Name())
			}
		}

		if dirBool && !*nonRecurs {
			runSearch(directory+file.Name(), search)
		}

	}

}

func fetchFiles(directory string) []os.FileInfo {
	var files []os.FileInfo
	var err error

	files, err = ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func dirTest(filename string) bool {

	fileStat, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Here's your problem")
		log.Fatal(err)
	}

	fileMode := fileStat.Mode()

	if fileMode.IsDir() {
		return true
	}

	return false

}
