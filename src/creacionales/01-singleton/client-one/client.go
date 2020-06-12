package client_one

import "github.com/lucaspichi06/patrones-diseno-go/creacionales/01-singleton/singleton"

// IncrementeAge incrementa la edad de la instancia de person
func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
