// B"H

/*
go mod init sandbox/csv-output
go run csv-output.go
*/

package main

import (
	"fmt"
	"os"
)

func removeDupes(s []string) []string {
	// -- -------------------------------
	//Make an empty map
	var m map[string]int
	m = make(map[string]int)

	// -- -------------------------------
	// Loop through each item in s
	for _, item := range s {

		// item's value increased by one. This adds it to the map if it wasn't there before
		m[item]++

	}

	// -- -------------------------------
	// List duplicates

	// for item, dupe := range m {

	// 	if dupe > 1 {
	// 		fmt.Println(item, "was entered", dupe, "times.")
	// 	}

	// }

	// -- -------------------------------
	// Make a list of unique values
	var result []string

	for item, _ := range m {
		result = append(result, item)
	}

	return result

}

func main() {

	// Get the user's list
	var testSlice []string = os.Args[1:]

	//Remove duplicates
	var resultSlice []string = removeDupes(testSlice)

	//Print results
	for _, item := range resultSlice {
		fmt.Printf("%s,", item)
	}
}
