package domain

type VehicleType int
type SlotType int

const (
	CarType VehicleType = iota
	BikeType
	TruckType
)
const (
	CarSlot SlotType = iota
	BikeSlot
	TruckSlot
)

type Vehicle interface {
	GetNo() int
	GetType() VehicleType
	GetSlot() SlotType
}

type Car struct {
	number      int
	vehicleType VehicleType
}

func NewCar(n int, vt VehicleType) *Car {
	return &Car{
		number:      n,
		vehicleType: vt,
	}
}

func (c *Car) GetNo() int {
	return c.number
}

func (c *Car) GetType() VehicleType {
	return c.vehicleType
}

func (c *Car) GetSlot() SlotType {
	return CarSlot

}

type Bike struct {
	number      int
	vehicleType VehicleType
}

func NewBike(n int, vt VehicleType) *Bike {
	return &Bike{
		number:      n,
		vehicleType: vt,
	}
}
func (c *Bike) GetNo() int {
	return c.number
}

func (c *Bike) GetType() VehicleType {
	return c.vehicleType
}

func (c *Bike) GetSlot() SlotType {
	return BikeSlot

}
