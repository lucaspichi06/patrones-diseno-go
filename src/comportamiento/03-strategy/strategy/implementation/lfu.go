package implementation

import (
	"fmt"

	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/03-strategy/strategy"
)

type Lfu struct{}

func (l *Lfu) Evict(c *strategy.Cache) {
	fmt.Println("Evicting by LFU strategy")
}
