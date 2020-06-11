// B''H

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	directory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileStat, err := os.Lstat(directory)
	if err != nil {
		fmt.Println("dirTest error")
		log.Fatal(err)
	}

	fileMode := fileStat.Mode()

	fmt.Println(fileMode.IsDir())

}
