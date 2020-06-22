package decorator

import "fmt"

type LogRegistry struct {
	Handler Decorator
}

// NewLogRegistry devuelve un decorator porque tiene que devolver un decorador para implementar la funcionalidad.-
func NewLogRegistry(handler Decorator) *LogRegistry {
	return &LogRegistry{handler}
}

func (lr *LogRegistry) Process() error {
	// fmt.Println("Petición") // todos los handler que quiera decorar primero van a ejecutar esto y después ejecutan su proceso.-
	defer fmt.Println("Petición") // si quiero ejecutar lo de este handler al final de lo que ejecuta el principal.-
	return lr.Handler.Process()
}
