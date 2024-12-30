package main

import (
	"poke/internal"
	"time"
)




func main() {
	interval := time.Duration(100*time.Second)
	clientPoke := internal.NewClient()
	cache := internal.NewCache(interval)
	cfg := &config{
		client: clientPoke,
		cache: *cache,
	}
	start(cfg)

	
}

