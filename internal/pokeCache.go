package internal

import (
	"encoding/json"
	"fmt"
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
}


func NewCache(interval time.Duration) *PokeCache{
	m := make(map[string]cacheEntry)
	return &PokeCache{
		entry: m,
	}
	
}

func (p *PokeCache)Add(key string, val []byte) {

	
	entry := cacheEntry{createdAt: time.Now(), value: val}
	
	if _, ok := p.entry[key]; !ok{
		p.entry[key] = entry
	}

	

}


func (p *PokeCache)Get(key string)([]byte, bool){
	if c, ok  := p.entry[key]; ok {
		return c.value, true
	}

	return nil, false
}

func (p *PokeCache) CacheClean(){
	fmt.Println("Cleaning Cache")
	keys := make([]string, 0, len(p.entry))
	for k := range p.entry {
		keys = append(keys, k)
	}
	
	for _, key := range keys {
		delete(p.entry,key)
		fmt.Printf("Removed %v from cache",key)
	}
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