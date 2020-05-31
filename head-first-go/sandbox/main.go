package main

import "fmt"

func newCar() *int64 {
	var carHonda int64
	carHonda = 10

	fmt.Println(carHonda)
	fmt.Println(&carHonda)

	return &carHonda
}

func main() {
	var carIntP *int64

	carIntP = newCar()

	*carIntP = 15

	fmt.Println(*carIntP)
	fmt.Println(carIntP)
}
