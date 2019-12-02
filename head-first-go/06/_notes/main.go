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

	// -- -------------------------
	fmt.Println("---------------")
	// -- -------------------------

	var office []string

	// Not needed if using append():
	// office = make([]string, 3)

	office = append(office, "computer")
	office = append(office, "desk")
	office = append(office, "pens")

	for index, value := range office {
		fmt.Println("index is", index, "- value is", value)
	}

	// With just println
	fmt.Println("Println:")
	fmt.Println(office)

	// -- -------------------------
	fmt.Println("---------------")
	// -- -------------------------

	fmt.Println("------------------------")
	fmt.Println("Slice example 3:")
	fmt.Println("Slice Literals")
	fmt.Println("No need to use `make`")
	fmt.Println("Use this 2% of the time")
	fmt.Println("--                    --")

	cities := []string{"b7", "jlm", "tlv", "tzfat", "elad"}

	cities[1] = "yerushalayim"

	for index, value := range cities {
		fmt.Println("index is", index, "- value is", value)
	}

	// -- -------------------------
	// General appropriate usage of slice
	fmt.Println("------------------------")

	shoppingList := make([]string, 7)

	fmt.Println("\n print 1:")
	fmt.Println(shoppingList)

	// -- -   -   -   -   -   -   -   -   -
	shoppingList[2] = "apples"
	shoppingList[4] = "bread"

	fmt.Println("\n print 2:")
	fmt.Println(shoppingList)

	// -- -   -   -   -   -   -   -   -   -
	fmt.Println("ALWAYS use same var name when appending!!!")
	shoppingList = append(shoppingList, "milk", "tomatoes", "wine")

	fmt.Println("\n print 3:")
	fmt.Println(shoppingList)

	// -- ---------------------------------

}
