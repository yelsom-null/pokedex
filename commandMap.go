package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func commandMap(cfg *config) error {
	
	url := "https://pokeapi.co/api/v2" + "/location-area"
	

	if cfg.nextLocationUrl != nil {
		url = *cfg.nextLocationUrl
	}


	
	if _, exists := cfg.cache.Get(url); !exists {
		fmt.Println("Cache not found")
	}

	

	


	locations, _ := cfg.client.GetLocations(cfg.nextLocationUrl)
	data, _ := json.Marshal(locations)
	cfg.cache.Add(strings.TrimSpace(url),data)
	// fmt.Printf("MAP: %v\n",url)
	
	

	cfg.nextLocationUrl = locations.Next
	cfg.previousLocationUrl = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
