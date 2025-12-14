package domain

import (
	"fmt"
)

type ParkingSlot struct {
	id       int
	level    int
	slotType SlotType
	vehicle  Vehicle
	isAvail  bool
}

var ParkingSlotId = 0

func GetId() int {
	ParkingSlotId++
	return ParkingSlotId
}
func NewParkingSlot(level int, st SlotType, v Vehicle, isAvail bool) *ParkingSlot {
	id := GetId()
	return &ParkingSlot{
		id:       id,
		level:    level,
		slotType: st,
		vehicle:  v,
		isAvail:  isAvail,
	}
}

func (p *ParkingSlot) IsAvailable() bool {
	return p.isAvail
}

func (p *ParkingSlot) GetSlotType() SlotType {
	return p.slotType
}

func (p *ParkingSlot) GetVehicle() (Vehicle, error) {
	if p.vehicle == nil {
		return nil, fmt.Errorf("No Vehicle in the parking slot with id :%v", p.id)
	}
	return p.vehicle, nil
}

func (p *ParkingSlot) GetId() int {
	return p.id
}

func (p *ParkingSlot) Park(v Vehicle) {

	p.vehicle = v
	p.isAvail = false

}

func (p *ParkingSlot) Unpark() {
	p.vehicle = nil
	p.isAvail = true
}
