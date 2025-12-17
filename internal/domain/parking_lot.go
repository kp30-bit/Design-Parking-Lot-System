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

func (pl *ParkingLot) Park(v Vehicle, parkingSlot *ParkingSlot) (*Ticket, error) {

	parkingSlot.Park(v)
	return &Ticket{
		ticket_id:      GetNewTicketId(),
		parkingSlot_id: parkingSlot.GetId(),
		vehicleNo:      v.GetNo(),
		price:          (int(v.GetType()) + 1) * 10, // base price based on the vehicle
		level:          parkingSlot.level,
		slotType:       parkingSlot.GetSlotType(),
	}, nil
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

func (pl *ParkingLot) GetParkingFloorMap() map[int]*ParkingFloor {
	tempMap := make(map[int]*ParkingFloor)
	for k, v := range pl.floors {
		tempMap[k] = v
	}
	return tempMap
}
