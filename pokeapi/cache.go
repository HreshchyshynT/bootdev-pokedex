package pokeapi

type Cache struct {
	storage map[string][]byte
}

func (c *Cache) Get(key string) ([]byte, bool) {
	v, ok := c.storage[key]
	return v, ok
}
func (c *Cache) Put(key string, value []byte) error {
	c.storage[key] = value
	return nil
}

func NewCache() *Cache {
	return &Cache{
		storage: make(map[string][]byte),
	}
}
