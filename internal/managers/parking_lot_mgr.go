package managers

import (
	"parking-lot-system/internal/domain"
	"sync"
)

type ParkingLotMgr struct {
	ParkingLot *domain.ParkingLot
}

var (
	ParkingLotMgrInstance *ParkingLotMgr
	singletonInstance     sync.Once
)

func NewParkingLotMgr(pl *domain.ParkingLot) *ParkingLotMgr {
	singletonInstance.Do(func() {
		ParkingLotMgrInstance = &ParkingLotMgr{
			ParkingLot: pl,
		}
	})
	return ParkingLotMgrInstance
}

func (p *ParkingLotMgr) Park(v domain.Vehicle) (*domain.Ticket, error) {
	ticket, err := p.ParkingLot.Park(v)
	if err != nil {
		return nil, err
	}
	return ticket, err
}

func (p *ParkingLotMgr) Unpark(ticket *domain.Ticket) {
	p.ParkingLot.UnPark(ticket)
}
