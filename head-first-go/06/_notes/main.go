package main

import "fmt"

func main() {

	// -- -------------------------
	// Look at Ylazerson's notes in this same place

	// Not necessary to do it this way. Sruly said about 8% of the time
	var hello []string
	hello = make([]string, 7)

	hello[2] = "hi"
	hello[3] = "how"
	hello[4] = "are"
	hello[5] = "you"

	for index, value := range hello {
		fmt.Println("index is", index, "- value is", value)
	}

	// -- -------------------------
	fmt.Println("---------------")
	// -- -------------------------
	// Do it like this 90% of the time
	names := make([]string, 7)

	names[3] = "Sruly"
	names[4] = "Yossi"
	names[5] = "Mendy"

	for index, value := range names {
		fmt.Println("index is", index, "- value is", value)
	}

}
