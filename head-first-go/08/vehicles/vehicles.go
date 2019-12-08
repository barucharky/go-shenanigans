package vehicles

type Brand struct {
	Make  string
	Model string
	Year  int
}

type Car struct {
	Brand
	Engine      string
	Mpg         float64
	PwrSteering bool
	Doors       int
}

type Truck struct {
	Brand
	Engine  string
	Fuel    string
	Wheels  int
	License bool
}

type Owner struct {
	Name string
	Cars []Car
}
