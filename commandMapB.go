package main

import (
	"encoding/json"
	"fmt"
	"poke/internal"
	
)


func commandMapB(cfg *config) error {
	
	url := "https://pokeapi.co/api/v2" + "/location-area"
	var locationsResp internal.Poke

	if cfg.previousLocationUrl != nil {
		url = *cfg.previousLocationUrl
	}

	
	
	if data, exists := cfg.cache.Get(url); exists {
		
	err := json.Unmarshal(data, &locationsResp)
		if err != nil {
		 	fmt.Printf("unable to marshal %v",err)
			}
		fmt.Println("Found cache")
		fmt.Printf("MAPB: %v\n",url)
		cfg.nextLocationUrl = locationsResp.Next
		cfg.previousLocationUrl = locationsResp.Previous
		for _, loc := range locationsResp.Results{
			fmt.Println(loc.Name)
		}	
		return nil
	}


	


	locations, _ := cfg.client.GetLocations(cfg.previousLocationUrl)

	cfg.nextLocationUrl = locations.Next
	cfg.previousLocationUrl = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
