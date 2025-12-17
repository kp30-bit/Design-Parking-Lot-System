package domain

type ParkingStrategy int

const (
	ClosestAvailableParking ParkingStrategy = iota
	RandomAvailableParking
)
