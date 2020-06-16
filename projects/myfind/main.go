// B''H

/*
go mod init sandbox/00-sandbox
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
	"path/filepath"
	"regexp"
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

	if len(flag.Args()) == 0 {
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
		directory, err = filepath.Abs(*dir)
		if err != nil {
			log.Fatal(err)
		}
	}

	// -- ---------------------------------
	// Run the search

	runSearch(directory, search)

	// Print results
	if len(results) == 0 {
		fmt.Println("No matches")
	} else {
		fmt.Println("Search results:")
		for _, result := range results {
			fmt.Println(result)
		}
	}

}

func runSearch(directory, search string) {

	// Get the list of files in the search directory, make search lower case, declare dirBool
	files := fetchFiles(directory)
	search = strings.ToLower(search)
	var dirBool bool

	for _, file := range files {

		// declare fullPath, define filename
		var fullPath string = directory + "/" + file.Name()
		var filename string = strings.ToLower(file.Name())
		dirBool = dirTest(fullPath)

		// Check match
		match, err := regexp.MatchString(search, filename)
		if err != nil {
			log.Fatal(err)
		}

		// If it's a match, check if the file only or directory only flags were selected and append the results
		if match {
			if *fileOnly && !dirBool {
				results = append(results, fullPath)
			} else if *dirOnly && dirBool {
				results = append(results, fullPath)
			} else if !*fileOnly && !*dirOnly {
				results = append(results, fullPath)
			}
		}

		// If the file is a directory, search within it as well
		if dirBool && !*nonRecurs {
			runSearch(fullPath, search)
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

	fileStat, err := os.Lstat(filename)
	if err != nil {
		log.Fatal(err)
	}

	fileMode := fileStat.Mode()

	return fileMode.IsDir()

}
