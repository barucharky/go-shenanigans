// B"H

package main

import (
	"fmt"

	"github.com/barucharky/go-shenanigans/head-first-go/08/magazine"
)

func main() {

	// -- -----------------------------------------------
	fmt.Println("Create subscriber1")

	subscriber1 := magazine.Subscriber{
		Name: "Aman Singh",
		Rate: 67.4,
	}

	// --   -   -   -   -   -   -   -   -   -   -   -   -
	fmt.Println("\n subscriber1 detail round 1")
	fmt.Println(subscriber1)

	// --   -   -   -   -   -   -   -   -   -   -   -   -
	subscriber1.Active = true
	subscriber1.Street = "123 Oak St"
	subscriber1.State = "NE"
	subscriber1.PostalCode = "68111"

	// --   -   -   -   -   -   -   -   -   -   -   -   -
	fmt.Println("\n subscriber1 detail round 2")
	fmt.Println(subscriber1)

	fmt.Println("Street:", subscriber1.Street)

	// -- -----------------------------------------------
	fmt.Println("\n Create subscriber2")

	var subscriber2 magazine.Subscriber

	subscriber2.Name = "Rasha B'galui"

	// --   -   -   -   -   -   -   -   -   -   -   -   -
	fmt.Println("\n subscriber2 detail round 1")
	fmt.Println(subscriber2)

	subscriber2.Rate = 99.87
	subscriber2.Street = "999 Main St"
	subscriber2.City = "Miami"
	subscriber2.State = "FL"

	// --   -   -   -   -   -   -   -   -   -   -   -   -
	fmt.Println("\n subscriber2 detail round 2")
	fmt.Println(subscriber2)

	fmt.Println("Street:", subscriber2.Street)
	fmt.Println("City:", subscriber2.City)

	// -- -----------------------------------------------

	// -- -----------------------------------------------
	fmt.Println("\n Create employee")
	employee := magazine.Employee{Name: "Joy Carr"}

	employee.Street = "456 Elm St"
	employee.City = "Portland"
	employee.State = "OR"
	employee.PostalCode = "97222"

	fmt.Println("\n employee detail")
	fmt.Println(employee)

	// -- -----------------------------------------------
	fmt.Println("\n Create addr")

	addr := magazine.Address{Street: "174th stg"}

	addr.City = "Kan tzivah Hashem es haBrachah"
	addr.State = "NY"
	addr.PostalCode = "11213"

	fmt.Println("\n addr detail")
	fmt.Println(addr)

	fmt.Println("Street:", addr.Street)
	fmt.Println("City:", addr.City)
	fmt.Println("State:", addr.State)
	fmt.Println("Postal Code:", addr.PostalCode)

	fmt.Println("\n add addr to employee")
	employee.Address = addr
	fmt.Println(employee)

	// -- -----------------------------------------------
	fmt.Println("\n Create some magazines")

	// --  -  -  -  -  -  -  -  -  -
	mag1 := magazine.Magazine{
		MagazineTitle: "Moshiach Time",
		Price:         5.56,
	}

	// --  -  -  -  -  -  -  -  -  -
	mag2 := magazine.Magazine{
		MagazineTitle: "Chayenu",
		Price:         1.99,
	}

	// --  -  -  -  -  -  -  -  -  -
	var mags []magazine.Magazine
	mags = append(mags, mag1)
	mags = append(mags, mag2)

	fmt.Println("mags variable:", mags)

	// --  -  -  -  -  -  -  -  -  -
	subscriber1.Magazines = mags

	fmt.Println(subscriber1)

}
