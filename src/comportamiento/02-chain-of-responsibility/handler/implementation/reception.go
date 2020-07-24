package implementation

import (
	"fmt"

	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/02-chain-of-responsibility/domain"
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/02-chain-of-responsibility/handler"
)

type Reception struct {
	next handler.Department
}

func (r *Reception) Execute(p *domain.Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
	}
	fmt.Println("Reception registrering patient")
	p.RegistrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next handler.Department) {
	r.next = next
}
