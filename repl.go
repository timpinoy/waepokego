package main

import (
	"bufio"
	"fmt"
	"github.com/timpinoy/waepokego/internal/pokeapi"
	"os"
)

const basePokeAPIUrl string = "https://pokeapi.co/api/v2/"

func NewConfig(client *pokeapi.Client) cliCommandConfig {
	return cliCommandConfig{
		Client: client,
	}
}

func runREPL(config cliCommandConfig) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		if command, ok := getCliCommands()[userInput]; ok {
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Print("\nUnknown command\n\n")
		}

		if scanner.Err() != nil {
			fmt.Println("scan error")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *cliCommandConfig) error
}

type cliCommandConfig struct {
	Previous *string
	Next     *string
	Client   *pokeapi.Client
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next set of map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous set of map locations",
			callback:    commandMapb,
		},
	}
}
