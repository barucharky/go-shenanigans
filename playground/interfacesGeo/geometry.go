// B''H

/*
go mod init sandbox/geometry
go run geometry.go
*/
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
)

var width *float64 = flag.Float64("w", 0, "width of a rectangle")
var length *float64 = flag.Float64("l", 0, "length of a rectangle")
var radius *float64 = flag.Float64("r", 0, "radius of a circle")

var errChoose error = errors.New("cannot calculate rectangle and circle")
var errNeg error = errors.New("only positive and non-zero numbers")

//step1) define our interface and input the methods which the interface will be comprised of:
type geometry interface {
	area() float64
	perim() float64
}

//step2) define our structs in order to establish which shapes will be used in this program
type rectangle struct {
	width  float64
	length float64
}
type circle struct {
	radius float64
}

//step 3) define our area methods which we will be using for each of our rectangle and circle structs
func (rec rectangle) area() float64 {
	return rec.width * rec.length
}
func (circ circle) area() float64 {
	return math.Pi * circ.radius * circ.radius
}

//step 4) define our perimeter methods which we will be using for each of our rectangle and circle structs
func (rec rectangle) perim() float64 {
	return (rec.width * 2) + (rec.length * 2)
}
func (circ circle) perim() float64 {
	return math.Pi * circ.radius * 2
}

//step 7: define a new method which will pass in our geometry interface
func measure(g geometry) {
	//the purpose of this method is to print 3 things:
	//when we later pass in our recA struct variable, the 2 fields of width and length will be passed into g
	fmt.Println(g)
	//our interface includes both our area() methods. So it will know to refer to either of our area methods.
	//if we pass recA into g, it will use the 'rectangle' struct to know to refer to the rectangle area() method
	//if we pass CircA into g, it will use the 'circle' struct to know to refer to the circle area() method
	fmt.Printf("Area is %.2f cm2\n", g.area())
	//our interface includes both our perim() methods. So it will know to refer to either of our perim methods.
	//if we pass recA into g, it will use the 'rectangle' struct to know to refer to the rectangle perim() method
	//if we pass CircA into g, it will use the 'circle' struct to know to refer to the circle perim() method
	fmt.Printf("Perimeter is %.2f cm\n", g.perim())
}
func main() {

	// parse the arguments
	flag.Parse()

	var rec rectangle
	var circ circle

	// make sure the user didn't mix rectangles and circles, handle arguments
	if (*length != 0 && *radius != 0) || (*width != 0 && *radius != 0) {

		log.Fatal(errChoose)

	} else if *width < 0 || *length < 0 || *radius < 0 {

		log.Fatal(errNeg)

	} else if *width == 0 && *length == 0 && *radius == 0 {

		log.Fatal(errNeg)

	} else if *width > 0 && *length > 0 {

		rec.width = *width
		rec.length = *length
		measure(rec)

	} else if *radius > 0 {

		circ.radius = *radius
		measure(circ)

	}
}
