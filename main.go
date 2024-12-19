package main

import (
	"bufio"
	"encoding/json"

	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)





type Client struct {
	httpClient http.Client
}

type config struct {
	client Client
	nextLocationUrl     *string 
	previousLocationUrl *string 
}

type Poke struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func main() {


	clientPoke := NewClient()

	cfg := &config{
		client: clientPoke,
	}
	startProgram(cfg)
	

}

func startProgram(cfg *config){
	
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommand()[commandName]

		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}

	}


	
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommand() map[string] cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Help messages",
			callback: commandHelp,
		},

		"map":{
			name: "map",
			description: "Get next location",
			callback: commandMap,
		},

		"exit":{
			name: "exit",
			description: "Ends program",
			callback: commandExit,
		},
	}
}

func commandMap(cfg *config) error {
	locations, _ := cfg.client.getLocations(cfg.nextLocationUrl)

	cfg.nextLocationUrl = locations.Next
	cfg.previousLocationUrl = locations.Previous

	for _, loc := range locations.Results{
		fmt.Println(loc.Name)
	}
	
	return nil
}



func commandExit(cfg *config) error {

	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {

	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	return nil
}


func NewClient() Client {
	return Client{
		httpClient: http.Client{},
	}
	
}


func (c *Client)getLocations(pageUrl *string) (Poke, error) {
	url := "https://pokeapi.co/api/v2"  + "/location-area"

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

	var locationsResp Poke
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return Poke{}, err
	}

	return locationsResp, nil
}