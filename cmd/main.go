package main

import (
	"parking-lot-system/internal/domain"
	"parking-lot-system/internal/managers"
)

func main() {
	//intial setup
	car := domain.NewCar(123, domain.CarType)
	bike := domain.NewBike(456, domain.BikeType)

	parkingSlot1 := domain.NewParkingSlot(1, domain.BikeSlot, nil, true)
	parkingSlot2 := domain.NewParkingSlot(1, domain.CarSlot, nil, true)
	parkingSlot3 := domain.NewParkingSlot(2, domain.CarSlot, nil, true)
	parkingSlot4 := domain.NewParkingSlot(2, domain.BikeSlot, nil, true)

	level1 := domain.NewParkingFloor(1)
	level2 := domain.NewParkingFloor(2)

	parkingLot := domain.NewParkingLot()

	parkingLot.AddFloor(level1)
	parkingLot.AddFloor(level2)

	level1.AddParkingSlot(parkingSlot1)
	level1.AddParkingSlot(parkingSlot2)
	level2.AddParkingSlot(parkingSlot3)
	level2.AddParkingSlot(parkingSlot4)

	ticketMgr := managers.NewTicketMgr()
	strategyMgr := managers.NewStrategyMgr()
	parkingLotMgr := managers.NewParkingLotMgr(parkingLot, strategyMgr, ticketMgr)
	ticket, _ := parkingLotMgr.Park(car, domain.ClosestAvailableParking)
	parkingLotMgr.Unpark(ticket)
	parkingLotMgr.Park(bike, domain.ClosestAvailableParking)
	parkingLotMgr.ParkingLot.ShowAllParkedVehicles()

}
