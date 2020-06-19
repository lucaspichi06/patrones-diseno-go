package todo

import "github.com/lucaspichi06/patrones-diseno-go/src/estructurales/07-bridge/list"

// Any recibe tareas repetidas.-
type Any struct {
	rendering list.List
	todos     []string
}

func NewAny() *Any {
	return &Any{}
}

func (a *Any) SetList(l list.List) {
	a.rendering = l
}

// Add agrega una tarea a la lista.-
func (a *Any) Add(name string) {
	a.todos = append(a.todos, name)
}

//Print muestra la lista con su respectiva representación.-
func (a *Any) Print() {
	a.rendering.Print(a.todos)
}
