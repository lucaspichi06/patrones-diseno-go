package implementation

import (
	"fmt"

	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/02-chain-of-responsibility/domain"
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/02-chain-of-responsibility/handler"
)

type Doctor struct {
	next handler.Department
}

func (d *Doctor) Execute(p *domain.Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
	}
	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true

	if d.next != nil {
		d.next.Execute(p)
	}
}

func (d *Doctor) SetNext(next handler.Department) {
	d.next = next
}
