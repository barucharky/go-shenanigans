// B''H

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	// get search
	var search string = os.Args[1]
	directory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Get the list of files in the search directory
	files := fetchFiles(directory)
	search = strings.ToLower(search)

	for _, file := range files {

		var filename string = strings.ToLower(file.Name())

		match, err := regexp.MatchString(search, filename)
		if err != nil {
			log.Fatal(err)
		}

		if match {
			fmt.Println("Match:\n", filename)
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
