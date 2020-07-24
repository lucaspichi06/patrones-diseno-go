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
	if c.next == nil {
		fmt.Println("Transaction Completed")
		return
	}
	if p.PaymentDone {
		fmt.Println("Payment already done")
		c.next.Execute(p)
	}
	fmt.Println("Payment in course")
	p.PaymentDone = true
	c.next.Execute(p)
}

func (c *Cashier) SetNext(next handler.Department) {
	c.next = next
}
