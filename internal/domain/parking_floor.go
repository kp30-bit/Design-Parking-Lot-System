package domain

type ParkingFloor struct {
	level int
	mp    map[int]*ParkingSlot
}

func NewParkingFloor(level int) *ParkingFloor {
	return &ParkingFloor{
		level: level,
		mp:    make(map[int]*ParkingSlot),
	}
}

func (pf *ParkingFloor) GetLevel() int {
	return pf.level
}

func (pf *ParkingFloor) GetAllSlots() map[int]*ParkingSlot {
	mapcopy := make(map[int]*ParkingSlot)
	for k, v := range pf.mp {
		mapcopy[k] = v
	}
	return mapcopy
}

func (pf *ParkingFloor) AddParkingSlot(ps *ParkingSlot) {
	pf.mp[ps.id] = ps
}

func (pf *ParkingFloor) GetParkingFloorMap() map[int]*ParkingSlot {
	tempMap := make(map[int]*ParkingSlot)
	for key, val := range pf.mp {
		tempMap[key] = val
	}
	return tempMap
}
