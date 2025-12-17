package interfaces

import "parking-lot-system/internal/domain"

type IParkingStrategy interface {
	GetFreeSlot(slotType domain.SlotType) (*domain.ParkingSlot, error)
}

//This interface enables strategy pattern.
