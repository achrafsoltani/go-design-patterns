package main

import "fmt"

func main() {
	fmt.Println("Singleton")
	for i := 0; i < 30; i++ {
		go GetInstance()
	}

	fmt.Println("Builder")
	manufacturingComplex := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()
	fmt.Println(car)
	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	motorbike := bikeBuilder.GetVehicle()
	fmt.Println(motorbike)
}
