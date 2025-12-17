package managers

import (
	"fmt"
	"parking-lot-system/internal/domain"
	"sync"
)

type ParkingLotMgr struct {
	ParkingLot  *domain.ParkingLot
	StrategyMgr *StrategyMgr
	TicketMgr   *TicketMgr
}

var (
	ParkingLotMgrInstance *ParkingLotMgr
	singletonInstance     sync.Once
)

func NewParkingLotMgr(pl *domain.ParkingLot, st *StrategyMgr, t *TicketMgr) *ParkingLotMgr {
	singletonInstance.Do(func() {
		ParkingLotMgrInstance = &ParkingLotMgr{
			ParkingLot:  pl,
			StrategyMgr: st,
			TicketMgr:   t,
		}
	})
	return ParkingLotMgrInstance
}

func (p *ParkingLotMgr) Park(v domain.Vehicle, strategy domain.ParkingStrategy) (*domain.Ticket, error) {
	parkingStrategy := p.StrategyMgr.SelectStrategy(strategy, p.ParkingLot.GetParkingFloorMap())
	parkingSlot, err := parkingStrategy.GetFreeSlot(v.GetSlot()) //Strategy Pattern is used here
	if err != nil {
		return nil, fmt.Errorf("Error getting free parking slot : %v\n", err)
	}
	ticket, err := p.ParkingLot.Park(v, parkingSlot)
	if err != nil {
		return nil, fmt.Errorf("Error Parking your car : %v\n", err)
	}
	p.TicketMgr.AddTicket(ticket)
	if err != nil {
		return nil, fmt.Errorf("Error adding ticket to ticket map: %v\n", err)
	}
	return ticket, nil
}

func (p *ParkingLotMgr) Unpark(ticket *domain.Ticket) {
	p.ParkingLot.UnPark(ticket)
}
