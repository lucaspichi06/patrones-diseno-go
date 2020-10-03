package main

import (
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/07-mediator/domain"
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/07-mediator/mediator"
)

func main() {
	stationManager := mediator.NewStationManager()
	passangerTrain := &domain.PassangerTrain{
		Mediator: stationManager,
	}
	goodsTrain := &domain.GoodsTrain{
		Mediator: stationManager,
	}
	passangerTrain.RequestArrival()
	goodsTrain.RequestArrival()
	passangerTrain.Departure()
}
