package singleton

import "sync"

var (
	p    *person
	once sync.Once
)

// GetInstance nos asegura el acceso único a la variable p
func GetInstance() *person {
	// Permite de manera segura que el proceso se ejecute una sola vez.-
	// Este método "Do" se ejecuta una única vez por paquete. Es decir que no puedo hacer dos veces once.Do().-
	once.Do(func() {
		p = &person{}
	})

	// Esta forma de implementar no permite controlar concurrencia.-
	// if p == nil {
	// 	p = &person{}
	// }
	return p
}

// person struct - se crea privada para que no se puedan crear "instancias" de esta estructura.-
type person struct {
	name string
	age  int
	//bloquea momentaneamente el proceso y asegura que se ejecute correctamente lo que necesito
	mux sync.Mutex
}

func (p *person) SetName(n string) {
	p.mux.Lock()
	p.name = n
	p.mux.Unlock()
}

func (p *person) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *person) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *person) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}
