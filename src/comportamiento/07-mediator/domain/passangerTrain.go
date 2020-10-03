package domain

import "fmt"

type PassangerTrain struct {
	Mediator mediator
}

func (p *PassangerTrain) RequestArrival() {
	if p.Mediator.CanLand(p) {
		fmt.Println("PassangerTrain: Landing")
	} else {
		fmt.Println("PassangerTrain: Waiting")
	}
}

func (p *PassangerTrain) Departure() {
	fmt.Println("PassangerTrain: Leaving")
	p.Mediator.NotifyFree()
}

func (p *PassangerTrain) PermitArrival() {
	fmt.Println("PassangerTrain: Arrival Permitted. Landing")
}
