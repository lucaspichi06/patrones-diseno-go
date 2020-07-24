package implementation

import (
	"fmt"

	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/02-chain-of-responsibility/domain"
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/02-chain-of-responsibility/handler"
)

type Medical struct {
	next handler.Department
}

func (m *Medical) Execute(p *domain.Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
	}
	fmt.Println("Medical giving medicine to patient")
	p.MedicineDone = true
	m.next.Execute(p)
}

func (m *Medical) SetNext(next handler.Department) {
	m.next = next
}
