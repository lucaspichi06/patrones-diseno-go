package strategy

type Cache struct {
	Storage      map[string]string
	EvictionAlgo EvictionAlgo
	Capacity     int
	MaxCapacity  int
}

func InitCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		Storage:      storage,
		EvictionAlgo: e,
		Capacity:     0,
		MaxCapacity:  2,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.EvictionAlgo = e
}

func (c *Cache) Add(key, value string) {
	if c.Capacity == c.MaxCapacity {
		c.evict()
	}
	c.Capacity++
	c.Storage[key] = value
}

func (c *Cache) Get(key string) {
	delete(c.Storage, key)
}

func (c *Cache) evict() {
	c.EvictionAlgo.Evict(c)
	c.Capacity--
}
