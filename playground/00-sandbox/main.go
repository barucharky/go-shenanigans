// B''H

/*
go mod init sandbox/sandbox
go run main.go
*/

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

	// Get the list of files in the search directory, make search lower case, declare dirBool
	files := fetchFiles("/home/baruch/bin")
	search := strings.ToLower("confirm")
	var dirBool bool

	for _, file := range files {

		// declare fullPath, define filename
		fmt.Println(file.Name())
		fullPath := "/home/baruch/bin/" + file.Name()

		fmt.Println("this is fullPath:", fullPath)

		var filename string = strings.ToLower(file.Name())
		dirBool = dirTest(fullPath)
		if dirBool {
			fmt.Println("It is a directory")
		}

		// Check match
		match, err := regexp.MatchString(search, filename)
		if err != nil {
			log.Fatal(err)
		}

		// If it's a match, check if the file only or directory only flags were selected and append the results
		if match {
			fmt.Println("Match:", fullPath)
		}

	}

}

func fetchFiles(directory string) []os.FileInfo {
	fmt.Println("In fetchFiles, this is directory:", directory)

	var files []os.FileInfo
	var err error

	files, err = ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("These are the files that fetchFiles returns:")
	for _, file := range files {
		fmt.Println(file.Name())
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
