package internal

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

/*

I used a Cache struct to hold a map[string]cacheEntry
and a mutex to protect the map across goroutines. A cacheEntry should be a struct with two fields:

createdAt - A time.Time that represents when the entry was created.
val - A []byte that represents the raw data we're caching.
You'll probably want to expose a NewCache() function that creates a new cache with a configurable interval (time.Duration).
*/

type cacheEntry struct {
	createdAt time.Time
	value []byte
	
}

type PokeCache struct {
	entry map[string]cacheEntry
	mu *sync.Mutex
}


func NewCache(interval time.Duration) *PokeCache{
	m := make(map[string]cacheEntry)

	p := &PokeCache{
		entry: m,
		mu:   &sync.Mutex{},
	}

	go p.reapLoop(interval)
	
	return p
}


func (p *PokeCache)reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	
	for range ticker.C{
		p.reap(time.Now().UTC(),interval)
	}
}

func (p *PokeCache)reap(now time.Time, last time.Duration){
	p.mu.Lock()
	defer p.mu.Unlock()
	for k, v := range p.entry {
		if v.createdAt.Before(now.Add(-last)) {
			fmt.Printf("\nDeleting: %v",k)
			delete(p.entry, k)
		}
	}
}

func (p *PokeCache)Add(key string, val []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	entry := cacheEntry{createdAt: time.Now(), value: val}
	
	if _, ok := p.entry[key]; !ok{
		p.entry[key] = entry
	}

	

}




func (p *PokeCache)Get(key string)([]byte, bool){
	p.mu.Lock()
	defer p.mu.Unlock()
	if c, ok  := p.entry[key]; ok {
		return c.value, true
	}

	return nil, false
}




func (p *PokeCache)DisplayCach(v []byte){
	var locationsResp Poke

	json.Unmarshal(v, &locationsResp)

	for _, mapEntry := range p.entry{
		fmt.Printf("Created: %v, Data: %v\n",mapEntry.createdAt,locationsResp.Results)
	}
}

func (p *PokeCache)ShowAll(){
	for _, e := range p.entry{

		fmt.Printf("Cache: %v\n", string(e.value))
		fmt.Println("-----------------------------")
	}	
}