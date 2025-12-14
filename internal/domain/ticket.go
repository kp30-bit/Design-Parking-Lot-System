package domain

type Ticket struct {
	ticket_id      int
	parkingSlot_id int
	vehicleNo      int
	price          int
	level          int
	slotType       SlotType
}

func NewTicket(tid, pid, no, price, level int, slotType SlotType) *Ticket {
	return &Ticket{
		ticket_id:      tid,
		parkingSlot_id: pid,
		price:          price,
		vehicleNo:      no,
		slotType:       slotType,
		level:          level,
	}
}

var (
	ticketId = 0
)

func GetNewTicketId() int {
	ticketId++
	return ticketId
}

func (t *Ticket) GetTicketId() int {
	return t.ticket_id
}

func (t *Ticket) GetParkingSlotId() int {
	return t.parkingSlot_id
}

func (t *Ticket) GetVehicle() int {
	return t.vehicleNo
}

func (t *Ticket) GetPrice() int {
	return t.price
}
func (t *Ticket) GetLevel() int {
	return t.level
}
