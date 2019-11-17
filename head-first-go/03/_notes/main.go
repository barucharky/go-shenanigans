// B''H

package main

import (
	"fmt"
)

func main() {

	// the % is a formatting verb
	fmt.Printf(
		"The %s cost %d cents in each\n",
		"gumballs",
		23,
	)
	fmt.Println("-------------------------")

	var houseCost float64 = 14.18
	var houses int64 = 515
	var name string = "Baruch"
	var isHappy bool = true

	// -- -----------------------------------------
	// Basic output
	fmt.Printf("One house costs %f \n", houseCost)
	fmt.Printf("There are %d houses \n", houses)
	fmt.Printf("The man's name is %s \n", name)
	fmt.Printf("The man's happiness is %t \n", isHappy)

	// -- -----------------------------------------
	// Basic output with set minimum width
	fmt.Println("-------------------------")
	fmt.Printf("One house costs %10f \n", houseCost)
	fmt.Printf("There are %10d houses \n", houses)
	fmt.Printf("The man's name is %10s \n", name)
	fmt.Printf("The man's happiness is %10t \n", isHappy)

	// -- -----------------------------------------
	// Basic output with set minimum width
	fmt.Println("-------------------------")
	fmt.Printf("%15s | %10d | %5.3f  \n",
		"Baruch",
		4685,
		1545.63,
	)
	fmt.Printf("%15s | %10d | %5.3f  \n",
		"Mendy",
		4115163221,
		0.5,
	)
	fmt.Printf("%15s | %10d | %5.3f  \n",
		"Yossi",
		8,
		324.48548,
	)
	fmt.Printf("%15s | %10d | %5.3f  \n",
		"Sruli",
		16,
		9.0,
	)

	fmt.Println("-------------------------")
	fmt.Println("Value and type of the variable, `name`")
	fmt.Printf("The value is %#v \n", name)
	fmt.Printf("The type is %T \n", name)

	fmt.Println("\nFunky characters:")
	var funkyChar string = "\t"
	fmt.Printf("z%sz \n", funkyChar)
	fmt.Printf("z%vz \n", funkyChar)
	fmt.Printf("z%#vz \n", funkyChar)
}
