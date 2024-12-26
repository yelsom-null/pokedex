package main

import "fmt"

func commandMap(cfg *config) error {
	locations, _ := cfg.client.GetLocations(cfg.nextLocationUrl)

	cfg.nextLocationUrl = locations.Next
	cfg.previousLocationUrl = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
