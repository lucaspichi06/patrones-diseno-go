package decorator

import "fmt"

type HandlerSayHello struct{}

func NewHandlerSayHello() *HandlerSayHello {
	return &HandlerSayHello{}
}

func (h *HandlerSayHello) Process() error {
	// fmt.Println("Petición") // esto no debe estar acá porque no es su funcionalidad.-
	fmt.Println("Hola")
	return nil
}
