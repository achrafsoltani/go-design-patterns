package main

import (
	"fmt"
	"log"
	"soltani.ma/go/design-patterns/creational"
)

func main() {
	// Singleton
	fmt.Println("Singleton")
	for i := 0; i < 30; i++ {
		go creational.GetInstance()
	}

	// Builder
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

	// Factory
	fmt.Println("Factory")
	payment, err := creational.GetPaymentMethod(creational.Cash)
	if err != nil {
		log.Fatal("Error getting payment method Cash")
	}
	msg := payment.Pay(10.30)
	fmt.Println(msg)
	payment, err = creational.GetPaymentMethod(creational.CreditCard)
	msg = payment.Pay(22.30)
	fmt.Println(msg)
}
