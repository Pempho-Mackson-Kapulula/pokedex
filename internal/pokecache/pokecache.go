package pokecache

import(
	"time"
	"sync"
)


type Cache struct{
	cache map[string]cacheEntry
	mu sync.Mutex
	interval time.Duration
}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

func NewCache (interval time.Duration) *Cache{
	c := Cache{
		cache: make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()

	return &c
}

func (c *Cache) Add(key string, val []byte){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}


func (c *Cache)Get(key string)([]byte,bool){
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheEntryStruct, ok := c.cache[key]

	if !ok{
		return nil, false

	}

	return cacheEntryStruct.val, true
}

func (c *Cache)reapLoop(){
	//create a ticker
	ticker := time.NewTicker(c.interval)
	 
	for range ticker.C{
		c.mu.Lock()
		for key , cacheEntryStruct := range c.cache{
			entryAge := time.Since(cacheEntryStruct.createdAt)
			if entryAge > c.interval{
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}