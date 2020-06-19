package decorator

import "fmt"

type HandlerSayGoodbye struct{}

func NewHandlerSayGoodbye() *HandlerSayGoodbye {
	return &HandlerSayGoodbye{}
}

func (h *HandlerSayGoodbye) Process() error {
	// fmt.Println("Petición") // esto no debe estar acá porque no es su funcionalidad.-
	fmt.Println("Adios")
	return nil
}
