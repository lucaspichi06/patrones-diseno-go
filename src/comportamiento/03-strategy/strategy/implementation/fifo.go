package implementation

import (
	"fmt"

	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/03-strategy/strategy"
)

type Fifo struct{}

func (l *Fifo) Evict(c *strategy.Cache) {
	fmt.Println("Evicting by FIFO strategy")
}
