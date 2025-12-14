package managers

import (
	"fmt"
	"parking-lot-system/internal/domain"
	"sync"
)

type TicketMgr struct {
	mp map[int]*domain.Ticket
}

var (
	TicketMgrInstance  *TicketMgr
	TicketMgrSingleton sync.Once
)

func NewTicketMgr() *TicketMgr {
	return &TicketMgr{
		mp: make(map[int]*domain.Ticket),
	}
}

func (tm *TicketMgr) AddTicket(ticket *domain.Ticket) {
	tm.mp[ticket.GetTicketId()] = ticket
}

func (tm *TicketMgr) ShowAllTickets() {
	for _, ticket := range tm.mp {
		fmt.Printf("Ticket no. %v|\tVehicle no. %v|\tPrice : %v|\t Slot Id: %v|\n", ticket.GetTicketId(), ticket.GetVehicle(), ticket.GetPrice(), ticket.GetParkingSlotId())
	}
}
