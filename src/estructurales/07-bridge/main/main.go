package main

import (
	"github.com/lucaspichi06/patrones-diseno-go/src/estructurales/07-bridge/list"
	"github.com/lucaspichi06/patrones-diseno-go/src/estructurales/07-bridge/todo"
)

func main() {
	myToDos := factortyToDo("unique")
	rendering := factoryList("bullet")
	myToDos.SetList(rendering)

	myToDos.Add("¿Qué estudiar?")
	myToDos.Add("Explicarlo con palabras sencillas")
	myToDos.Add("Hacer con lo que aprendimos")
	myToDos.Add("Revisar y mejorar")
	myToDos.Add("¿Qué estudiar?")
	myToDos.Print()
}

func factortyToDo(s string) todo.ToDo {
	if s == "unique" {
		return todo.NewUnique()
	}

	return todo.NewAny()
}

func factoryList(s string) list.List {
	if s == "plain" {
		return list.NewPlain()
	}

	return list.NewBullet('*')
}
