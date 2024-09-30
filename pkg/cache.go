// cache.go
package pkg

import (
	"sync"
)

// Cache represents a simple in-memory cache for storing instructions
type Cache struct {
	data map[uint32]uint32 // Maps instruction addresses to instruction values
	mu   sync.RWMutex      // Mutex for concurrent access
}

// NewCache creates a new Cache instance
func NewCache() *Cache {
	return &Cache{
		data: make(map[uint32]uint32),
	}
}

// Get retrieves an instruction from the cache
func (c *Cache) Get(address uint32) (uint32, bool) {
	c.mu.RLock() // Acquire read lock
	defer c.mu.RUnlock()
	instruction, found := c.data[address]
	return instruction, found
}

// Set stores an instruction in the cache
func (c *Cache) Set(address uint32, instruction uint32) {
	c.mu.Lock() // Acquire write lock
	defer c.mu.Unlock()
	c.data[address] = instruction
}

// Clear clears the cache
func (c *Cache) Clear() {
	c.mu.Lock() // Acquire write lock
	defer c.mu.Unlock()
	c.data = make(map[uint32]uint32) // Reset the cache
}
