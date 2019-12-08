package main

import (
	"fmt"

	"github.com/barucharky/go-shenanigans/head-first-go/08/vehicles"
)

func main() {

	fmt.Println("Cars and Trucks")
	fmt.Println("")

	fmt.Println("First Vehicle:")

	var truck1 vehicles.Truck

	truck1.Make = "Ford"
	truck1.Model = "F150"
	truck1.Year = 2011
	truck1.Engine = "V8"
	truck1.Fuel = "Gas"
	truck1.Wheels = 6
	truck1.License = false

	fmt.Println(truck1)

	fmt.Println("Make:", truck1.Make)
	fmt.Println("Model:", truck1.Model)
	fmt.Println("Year:", truck1.Year)
	fmt.Println("Engine:", truck1.Engine)
	fmt.Println("Fuel:", truck1.Fuel)
	fmt.Println("Wheels:", truck1.Wheels)
	fmt.Println("License:", truck1.License)

	fmt.Println("----")
	fmt.Println("Second Vehicle")

	var car1 vehicles.Car

	car1.Make = "Toyota"
	car1.Model = "celica"
	car1.Year = 2004
	car1.Doors = 2
	car1.Engine = "V6"
	car1.Mpg = 37.62
	car1.PwrSteering = true

	fmt.Println(car1)

	fmt.Println("Make:", car1.Make)
	fmt.Println("Model:", car1.Model)
	fmt.Println("Year:", car1.Year)
	fmt.Println("Doors:", car1.Doors)
	fmt.Println("Engine:", car1.Engine)
	fmt.Println("Mpg:", car1.Mpg)
	fmt.Println("PwrSteering:", car1.PwrSteering)

	fmt.Println("-----")
	fmt.Println("Third Vehicle")

	var car2 vehicles.Car

	car2.Make = "Mercedes"
	car2.Model = "Benz"
	car2.Year = 2019
	car2.Doors = 2
	car2.Engine = "V8"
	car2.Mpg = 21.38
	car2.PwrSteering = true

	fmt.Println(car2)

	fmt.Println("Make:", car2.Make)
	fmt.Println("Model:", car2.Model)
	fmt.Println("Year:", car2.Year)
	fmt.Println("Doors:", car2.Doors)
	fmt.Println("Engine:", car2.Engine)
	fmt.Println("Mpg:", car2.Mpg)
	fmt.Println("PwrSteering:", car2.PwrSteering)

	fmt.Println("-----")
	fmt.Println("Yossi's Cars")

	var cars []vehicles.Car

	cars = append(cars, car1)
	cars = append(cars, car2)

	var owner1 vehicles.Owner

	owner1.Name = "Yossi David"
	owner1.Cars = cars

	fmt.Println(owner1)

	fmt.Println("\nEach Car:")

	for _, car := range owner1.Cars {
		fmt.Println(car.Brand)
		fmt.Println("Doors:", car.Doors)
		fmt.Println("Engine:", car.Engine)
		fmt.Println("Mpg:", car.Mpg)
		fmt.Println("PwrSteering:", car.PwrSteering)
		fmt.Println("")
	}

}
