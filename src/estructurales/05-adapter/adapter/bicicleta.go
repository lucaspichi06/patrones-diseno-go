package adapter

import "github.com/lucaspichi06/patrones-diseno-go/src/estructurales/05-adapter/bici"

type BicicletaAdapter struct {
	Bicicleta *bici.Bicicleta
}

func (b *BicicletaAdapter) Mover() {
	b.Bicicleta.Avanzar()
}
