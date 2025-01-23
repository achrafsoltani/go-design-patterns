package creational

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	Builder BuildProcess
}

func (D *ManufacturingDirector) Construct() {
	D.Builder.SetWheels().SetSeats().SetStructure()
}
func (D *ManufacturingDirector) SetBuilder(builder BuildProcess) {
	D.Builder = builder
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (c *BikeBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}

func (c *BikeBuilder) SetSeats() BuildProcess {
	c.v.Seats = 2
	return c
}

func (c *BikeBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Bike"
	return c
}

func (c *BikeBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BusBuilder struct {
	v VehicleProduct
}

func (c *BusBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2 * 4
	return c
}

func (c *BusBuilder) SetSeats() BuildProcess {
	c.v.Seats = 30
	return c
}

func (c *BusBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Bus"
	return c
}

func (c *BusBuilder) GetVehicle() VehicleProduct {
	return c.v
}
