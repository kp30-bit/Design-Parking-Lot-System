package managers

import (
	"parking-lot-system/internal/domain"
	interfaces "parking-lot-system/internal/interface"
	"parking-lot-system/internal/usecase"
)

type StrategyMgr struct {
}

func NewStrategyMgr() *StrategyMgr {
	return &StrategyMgr{}
}

// using factory pattern here to create the required object based on input in runtime
func (sm *StrategyMgr) SelectStrategy(inputStrategy domain.ParkingStrategy, ParkingFloorMap map[int]*domain.ParkingFloor) interfaces.IParkingStrategy {
	switch inputStrategy {
	case domain.ClosestAvailableParking:
		return usecase.NewClosestAvailableParkingStrategy(ParkingFloorMap)
	}
	return nil
}
