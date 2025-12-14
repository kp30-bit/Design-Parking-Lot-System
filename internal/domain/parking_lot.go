package domain

import (
	"errors"
	"fmt"
)

type ParkingLot struct {
	floors map[int]*ParkingFloor
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{
		floors: make(map[int]*ParkingFloor),
	}
}

func (pl *ParkingLot) Park(v Vehicle) (*Ticket, error) {
	slotType := v.GetSlot()
	for _, parkingFloor := range pl.floors {
		parkingSlot := parkingFloor.GetFreeSlot(slotType)
		if parkingSlot != nil {
			parkingSlot.Park(v)
			return NewTicket(GetNewTicketId(), parkingSlot.GetId(), v.GetNo(), int(slotType)*10, parkingSlot.level, v.GetSlot()), nil
		}
	}
	return nil, errors.New("Couldn't find a parking for your vehicle\n")
}

func (pl *ParkingLot) UnPark(t *Ticket) error {
	parkingFloor := pl.floors[t.level]
	parkingSlotsMap := parkingFloor.mp

	parkingSlot, ok := parkingSlotsMap[t.parkingSlot_id]
	if ok {
		parkingSlot.Unpark()
		return nil
	}
	return errors.New("Invalid Ticket given \n")
}

func (pl *ParkingLot) AddFloor(pf *ParkingFloor) {
	pl.floors[pf.level] = pf
}

func (pl *ParkingLot) ShowAllParkedVehicles() {
	for _, floors := range pl.floors {
		for _, slots := range floors.mp {
			if !slots.isAvail {
				vehicle, _ := slots.GetVehicle()
				fmt.Printf("Vehicle no. %v is parked on parking slot no. %v on level %v\n", vehicle.GetNo(), slots.GetId(), slots.level)
			}
		}
	}
}
