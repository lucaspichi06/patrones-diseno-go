package todo

import "github.com/lucaspichi06/patrones-diseno-go/src/estructurales/07-bridge/list"

// ToDo interface que tiene la lista de tareas. Refier a la implementación.-
type ToDo interface {
	SetList(l list.List)
	Add(name string)
	Print()
}
