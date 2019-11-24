// B''H

package main

import "fmt"

func main() {

	// -- ---------------------------------
	// when values are not known in advance:
	var myInts [3]int64

	myInts[1] = 99

	fmt.Println("myInts:")
	fmt.Println(myInts[0])
	fmt.Println(myInts[1])
	fmt.Println(myInts[2])

	// -- ---------------------------------
	// when values are known in advance:
	myInts2 := [3]int64{23, 0, 99999}

	myInts2[2] = 45

	fmt.Println("myInts2:")
	fmt.Println(myInts2[0])
	fmt.Println(myInts2[1])
	fmt.Println(myInts2[2])

	// -- ---------------------------------
	// Multi-lines
	myStrings := [5]string{
		"Ayn",
		"Od",
		"Milvado",
		"Efes",
		"Zulaso",
	}

	fmt.Println("myStrings:")
	fmt.Println(myStrings[0])
	fmt.Println(myStrings[1])
	fmt.Println(myStrings[2])
	fmt.Println(myStrings[3])
	fmt.Println(myStrings[4])

	fmt.Println("myStrings default user display:")
	fmt.Printf("%v\n", myStrings)
	fmt.Println(`myStrings "go-code" display:`)
	fmt.Printf("%#v\n", myStrings)

	// -- ---------------------------------
	// Array loops
	fmt.Println("-------------")
	fmt.Println("Array loops 1")

	for index, value := range myStrings {
		fmt.Println("Index is", index, "- value is", value)
	}

	fmt.Println("-------------")
	fmt.Println("Array loops 2")

	for _, value := range myStrings {
		fmt.Println("value is", value)
	}

}
