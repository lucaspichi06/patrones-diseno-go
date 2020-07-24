package implementation

import (
	"fmt"

	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/02-chain-of-responsibility/domain"
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/02-chain-of-responsibility/handler"
)

type Cashier struct {
	next handler.Department
}

func (c *Cashier) Execute(p *domain.Patient) {
	if p.PaymentDone {
		fmt.Println("Payment already done")
		c.next.Execute(p)
	}
	fmt.Println("Cashier getting money from patient")
	p.PaymentDone = true
	if c.next != nil {
		c.next.Execute(p)
	}
}

func (c *Cashier) SetNext(next handler.Department) {
	c.next = next
}
