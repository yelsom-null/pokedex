package main

import (
	"bufio"
	"fmt"
	"os"
	"poke/internal"
	"strings"
)

type config struct {
	client              internal.Client
	nextLocationUrl     *string
	previousLocationUrl *string
}


type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func start(cfg *config) {

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

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Help messages",
			callback:    commandHelp,
		},

		"map": {
			name:        "map",
			description: "Get next location",
			callback:    commandMap,
		},

		"exit": {
			name:        "exit",
			description: "Ends program",
			callback:    commandExit,
		},
	}
}
