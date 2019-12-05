package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {

	// -- ---------------------------
	var directory string
	var err error

	directory, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(directory)
	// -- ---------------------------

	// Type os.FileMode
	fi, err := os.Lstat(directory)
	if err != nil {
		log.Fatal(err)
	}

	mode := fi.Mode()

	if mode.IsDir() {
		fmt.Println("It is a directory")
	}

	fmt.Println("-----")
	fmt.Println("mode is:", mode)
	fmt.Println("-----")

	fmt.Println("Variable types:")
	fmt.Println("fi is type:", reflect.TypeOf(fi))
	fmt.Println("mode is type:", reflect.TypeOf(mode))
	// -- ---------------------------

	// -- ---------------------------
	// Create a directory
	var startDir string = "/home/baruch/"
	var newDirName string = "stupid-dir/"

	fmt.Println("-----")
	fmt.Println("Creating directory in", startDir, "...")

	err = os.Chdir(startDir)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(newDirName, 0777)
	if err != nil {
		log.Fatal(err)
	}

	var endDir string = startDir + newDirName
	fmt.Println("Created", endDir)
	// -- ---------------------------

	fmt.Println("-----")
	fmt.Println("Creating file in", endDir)

	var fileName string = "pointless-file"
	var newFile *os.File

	newFile, err = os.Create(endDir + fileName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created", fileName)
	fmt.Println("newFile is", newFile)
	// -- ---------------------------

	fmt.Println("-----")
	fmt.Println("Writing to file...")

	var message string = "You are reading a pointless file.\n"

	_, err = newFile.WriteString(message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wrote a message in the file.")
	// -- ---------------------------

	fmt.Println("-----")
	fmt.Println("Checking filesize")

	var myInfo os.FileInfo

	myInfo, err = os.Stat(endDir + fileName)

	fmt.Println("File size is", myInfo.Size())
}
