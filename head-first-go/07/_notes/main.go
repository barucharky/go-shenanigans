// B"H

package main

import "fmt"

func main() {

	// -- ---------------------------------
	// Map example 1
	fmt.Println("------------------------")
	fmt.Println("Map example 1:")
	fmt.Println("Use this 90% of the time")

	shopList := make(map[string]int)

	shopList["shoes"] = 5
	shopList["hat"] = 1
	shopList["tie"] = 2
	shopList["cups"] = 50
	shopList["cups"] = 75

	fmt.Println(shopList)

	fmt.Println(shopList["cups"])
	fmt.Println(shopList["zzzz"])

	// -- _________________________________
	/*
		NOTE:
		    - A map is an unordered collection of keys and values.

			- When you use a for...range loop with a map, you never know what
		      order you’ll get the map’s contents in.
	*/
	fmt.Println("Map example 1 - loop:")

	for key, value := range shopList {
		fmt.Println("key is", key, "- value is", value)
	}
	// -- ---------------------------------

	// -- ---------------------------------
	// Map example 2
	fmt.Println("------------------------")
	fmt.Println("Map example 1, using map literals:")

	houses := map[string]float64{
		"21 main st":  78349.65,
		"943 loon st": 8989.65,
	}

	houses["943 loon st"] = 5.78
	houses["huh ave 99"] = 1.9898989

	fmt.Println(houses)
	// -- ---------------------------------

	// -- ---------------------------------
	fmt.Println("------------------------")
	fmt.Println("... having fun ...:")

	countries := make(map[string]int)

	countries["aus"]++
	countries["aus"]++

	fmt.Println(countries)
	// -- ---------------------------------

	// -- ---------------------------------
	fmt.Println("------------------------")
	fmt.Println("handling zero-vals correctly")

	var val int
	var ok bool

	val, ok = countries["aus"]

	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("does not exist in map")
	}

	// -- _________________________________

	val, ok = countries["usa"]

	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("does not exist in map")
	}

	// -- ---------------------------------

	delete(countries, "aus")

	val, ok = countries["aus"]

	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("does not exist in map")
	}

	// -- ---------------------------------

}
