package creational

import (
	"errors"
	"fmt"
)

type Vehicle interface {
	GetWheels() int
	GetSeats() int
}

type Car interface {
	GetDoors() int
}

type Motorbike interface {
	GetType() int
}

type VehicleFactory interface {
	GetVehicle(v int) (Vehicle, error)
}

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct{}

func (*CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("invalid vehicle type %d", v))
	}
}

type LuxuryCar struct{}

func (*LuxuryCar) GetDoors() int {
	return 2
}
func (*LuxuryCar) GetWheels() int {
	return 4
}
func (*LuxuryCar) GetSeats() int {
	return 2
}

type FamilyCar struct{}

func (*FamilyCar) GetDoors() int {
	return 5
}
func (*FamilyCar) GetWheels() int {
	return 4
}
func (*FamilyCar) GetSeats() int {
	return 5
}
