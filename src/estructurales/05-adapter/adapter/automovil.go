package adapter

import "github.com/lucaspichi06/patrones-diseno-go/src/estructurales/05-adapter/auto"

type AutomovilAdapter struct {
	Automovil *auto.Automovil
}

func (a *AutomovilAdapter) Mover() {
	a.Automovil.Encender()
	a.Automovil.Acelerar()
}