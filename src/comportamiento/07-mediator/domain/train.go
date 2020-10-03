package domain

type Train interface {
	RequestArrival()
	Departure()
	PermitArrival()
}