package adapter

import (
	"github.com/lucaspichi06/patrones-diseno-go/src/estructurales/05-adapter/auto"
	"github.com/lucaspichi06/patrones-diseno-go/src/estructurales/05-adapter/bici"
)

func Factory(s string) Adapter {
	switch s {
	case "automovil":
		d := auto.Automovil{}
		return &AutomovilAdapter{&d}
	case "bicicleta":
		d := bici.Bicicleta{}
		return &BicicletaAdapter{&d}
	}

	return nil
}
