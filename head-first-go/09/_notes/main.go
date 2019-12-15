// B''H

package main

import (
	"fmt"

	"github.com/barucharky/go-shenanigans/head-first-go/09/volumes"
)

func main() {

	var soda volumes.Liters
	var water volumes.Milliliters

	soda = 2
	water = 500

	// -- -------------------------
	fmt.Printf("Soda: %0.3f liters\n", soda)
	fmt.Printf("Water: %0.3f milliliters\n", water)

	// -- -------------------------
	var sodaGallons volumes.Gallons
	var waterGallons volumes.Gallons

	sodaGallons = soda.ToGallons()
	waterGallons = water.ToGallons()

	fmt.Println("-----")
	fmt.Printf("Gallons of soda: %0.3f\n", sodaGallons)
	fmt.Printf("Gallons of water: %0.3f\n", waterGallons)
	fmt.Printf("Total gallons of liquid: %0.3f\n", sodaGallons+waterGallons)

	// -- -------------------------
	// Methods with pointers
	fmt.Println("-----")

	soda.DoubleWrong()
	fmt.Println("Doubling without pointers:", soda)

	soda.Double()
	fmt.Println("Doubling with pointers:", soda)

}
