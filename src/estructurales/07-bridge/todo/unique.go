package todo

import "github.com/lucaspichi06/patrones-diseno-go/src/estructurales/07-bridge/list"

// Unique lista de tareas únicas.-
type Unique struct {
	rendering list.List
	todos     []string
}

func NewUnique() *Unique {
	return &Unique{}
}

func (u *Unique) SetList(l list.List) {
	u.rendering = l
}

// Add agrega una tarea a la lista.-
func (u *Unique) Add(name string) {
	for _, v := range u.todos {
		if name == v {
			return
		}
	}
	u.todos = append(u.todos, name)
}

//Print muestra la lista con su respectiva representación.-
func (u *Unique) Print() {
	u.rendering.Print(u.todos)
}
