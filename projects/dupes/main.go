package main

import (
	"flag"
	"fmt"
)

var n *bool = flag.Bool("n", false, "display non-duplicates")

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
	// Make a slice of unique values
	var result []string

	if *n {
		for item, occurrence := range m {
			if occurrence == 1 {
				result = append(result, item)
			}
		}
	} else {
		for item, occurrence := range m {
			if occurrence > 1 {
				result = append(result, item)
			}
		}
	}

	return result
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
