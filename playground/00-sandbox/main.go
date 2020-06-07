package main

import (
	"flag"
	"fmt"
)

var n *bool = flag.Bool("n", false, "display non-duplicates")
var was *string = flag.String("w", "was entered", "put indicator of multiples")

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
	for item, occurance := range m {

		if occurance > 1 {
			fmt.Println(item, *was, occurance, "times.")
		}

	}

	// -- -------------------------------
	// Make a slice of unique values
	var dupes []string
	var nonDupes []string

	for item, occurance := range m {

		if occurance > 1 {
			dupes = append(dupes, item)
		} else {
			nonDupes = append(nonDupes, item)
		}
	}

	if *n {
		return nonDupes
	}

	return dupes

}

func main() {

	flag.Parse()

	// Get the user's list
	var testSlice []string = flag.Args()

	//Remove duplicates
	var resultSlice []string = removeDupes(testSlice)

	//Print results
	fmt.Println(resultSlice)
}
