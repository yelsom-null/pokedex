package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

)

func (c *Client) GetLocations(pageUrl *string) (Poke, error) {
	url := "https://pokeapi.co/api/v2" + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Poke{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Poke{}, err
	}
	defer resp.Body.Close()



	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Poke{}, err
	}
	fmt.Printf("Adding %v to cache\n",url)


	var locationsResp Poke
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return Poke{}, err
	}

	return locationsResp, nil
}