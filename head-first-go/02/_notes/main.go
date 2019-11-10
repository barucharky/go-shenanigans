// B''H

package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now)

	var nowDay int = now.Day()
	fmt.Println(nowDay)

	var nowDay2 int = time.Now().Day()
	fmt.Println(nowDay2)

	// -- ---------------------------------
	var grade int = 65
	var passingGrade int = 60
	var b4If = "Before if"

	if grade >= passingGrade {
		fmt.Println("You passed!")
		fmt.Println(b4If)
	}

	var passed bool = grade <= passingGrade
	if passed {
		fmt.Println("You passed again")

		var congrats string = "You are doing great"
		fmt.Println(congrats)

	} else if !passed {
		fmt.Println("You failed")
	}
	// -- ---------------------------------

	// -- ---------------------------------
	fmt.Println("Starting for loop")

	for x := 1; x <= 6; x++ {
		fmt.Println(x)
	}

	fmt.Println("Starting 2nd for loop")
	for x := 10; x >= 6; x-- {
		fmt.Println(x)
	}

	fmt.Println("Starting 3rd for loop")
	for x := 10; x >= 6; x-- {

		fmt.Println("Before check")

		if x == 7 {
			continue
		}

		fmt.Println("After check")
		fmt.Println(x)
	}

	fmt.Println("Starting 3rd for loop")
	for x := 10; x >= 6; x-- {

		fmt.Println("Before check")

		if x == 7 {
			break
		}

		fmt.Println("After check")
		fmt.Println(x)
	}
	// -- ---------------------------------
}
