package main

import (
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/03-strategy/strategy"
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/03-strategy/strategy/implementation"
)

func main() {
	lfu := &implementation.Lfu{}
	cache := strategy.InitCache(lfu)
	cache.Add("a", "1")
	cache.Add("b", "2")
	cache.Add("c", "3")
	lru := &implementation.Lru{}
	cache.SetEvictionAlgo(lru)
	cache.Add("d", "4")
	fifo := &implementation.Fifo{}
	cache.SetEvictionAlgo(fifo)
	cache.Add("e", "5")
}
