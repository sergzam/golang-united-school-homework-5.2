package cache

import "time"

type Cache struct {
	values map[string]cachedValue
}

type cachedValue struct {
	value    string
	deadline time.Time
}

//goland:noinspection GoUnusedExportedFunction
func NewCache() Cache {
	return Cache{values: make(map[string]cachedValue)}
}

func (c Cache) Get(key string) (string, bool) {
	item, ok := c.values[key]
	if !ok {
		return "", false
	}

	if item.deadline.IsZero() || item.deadline.After(time.Now()) {
		return item.value, true
	}

	delete(c.values, key)
	return "", false
}

func (c Cache) Put(key, value string) {
	item := cachedValue{value: value}
	c.values[key] = item
}

func (c Cache) Keys() []string {
	keys := make([]string, 0, len(c.values))
	for k := range c.values {
		keys = append(keys, k)
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	item := cachedValue{value: value, deadline: deadline}
	c.values[key] = item
}
