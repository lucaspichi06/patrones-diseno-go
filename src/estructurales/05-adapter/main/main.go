package main

import (
	"fmt"

	"github.com/lucaspichi06/patrones-diseno-go/src/estructurales/05-adapter/adapter"
)

func main() {
	var t string
	fmt.Print("Digite el tipo de transporte: ")
	fmt.Scanln(&t)

	transportAdapter := adapter.Factory(t)
	transportAdapter.Mover()
}
