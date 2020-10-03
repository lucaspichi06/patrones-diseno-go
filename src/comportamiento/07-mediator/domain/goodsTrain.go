package domain

import "fmt"

type GoodsTrain struct {
	Mediator mediator
}

func (g *GoodsTrain) RequestArrival() {
	if g.Mediator.CanLand(g) {
		fmt.Println("GoodsTrain: Landing")
	} else {
		fmt.Println("GoodsTrain: Waiting")
	}
}

func (g *GoodsTrain) Departure() {
	fmt.Println("GoodsTrain: Leaving")
	g.Mediator.NotifyFree()
}

func (g *GoodsTrain) PermitArrival() {
	fmt.Println("GoodsTrain: Arrival Permitted. Landing")
}
