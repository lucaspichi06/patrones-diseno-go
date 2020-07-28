package implementation

import (
	"fmt"

	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/03-strategy/strategy"
)

type Lru struct{}

func (l *Lru) Evict(c *strategy.Cache) {
	fmt.Println("Evicting by LRU strategy")
}
