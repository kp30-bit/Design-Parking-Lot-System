package usecase

import (
	"fmt"
	"parking-lot-system/internal/domain"
)

type ClosestAvailableParkingStrategy struct {
	FloorMap map[int]*domain.ParkingFloor
}

func NewClosestAvailableParkingStrategy(mp map[int]*domain.ParkingFloor) *ClosestAvailableParkingStrategy {
	return &ClosestAvailableParkingStrategy{
		FloorMap: mp,
	}
}

func (c *ClosestAvailableParkingStrategy) GetFreeSlot(slotType domain.SlotType) (*domain.ParkingSlot, error) {
	if c.FloorMap == nil {
		return nil, fmt.Errorf("Floor Map not available")
	}
	for _, floor := range c.FloorMap {
		for _, slot := range floor.GetParkingFloorMap() {
			if slot.IsAvailable() && slot.GetSlotType() == slotType {
				return slot, nil
			}
		}
	}
	return nil, nil

}
