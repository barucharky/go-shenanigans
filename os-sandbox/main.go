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

	err = os.Chdir(startDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----")
	fmt.Println("Creating directory in", startDir)

	err = os.Mkdir("gotemp", 0777)
	if err != nil {
		log.Fatal(err)
	}

}
