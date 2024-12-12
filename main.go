package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


var cmds = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {

	ShowWelcome()
	ShowPrompt()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		switch input {
		case "exit":
			commandExit()
		case "help":
			commandHelp()
		default: fmt.Print("Unknown command")
		}
		

		words := strings.Fields(input)

		w := words[0]

		w = strings.ToLower(w)
		w = strings.TrimSpace(w)

		fmt.Printf("Your command was: %v\n\n", w)
		ShowPrompt()
	}

}

func ShowPrompt() {
	fmt.Print("Pokedex > ")
}

func ShowWelcome(){
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
}

func commandExit() error {

	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return nil
}

func commandHelp() error {
	
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

	for _,cmdVal := range cmds{
		fmt.Printf("%v\n\n",cmdVal.name)
	}
	return nil
}

