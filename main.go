package main

import (
	"fmt"
	"soltani.ma/go/design-patterns/creational"
)

func main() {
	fmt.Println("Singleton")
	for i := 0; i < 30; i++ {
		go creational.GetInstance()
	}

	fmt.Println("Builder")
	manufacturingComplex := creational.ManufacturingDirector{}
	carBuilder := &creational.CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()
	fmt.Println(car)
	bikeBuilder := &creational.BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	motorbike := bikeBuilder.GetVehicle()
	fmt.Println(motorbike)
}
