package main

import (
	"bufio"
	"fmt"
	"github.com/timpinoy/waepokego/internal/pokeapi"
	"github.com/timpinoy/waepokego/internal/pokedex"
	"os"
	"strings"
)

func NewConfig(client *pokeapi.Client, dex *pokedex.Pokedex) cliCommandConfig {
	return cliCommandConfig{
		Client:  client,
		Pokedex: dex,
	}
}

func runREPL(config cliCommandConfig) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		words := strings.Fields(strings.ToLower(userInput))
		if command, ok := getCliCommands()[words[0]]; ok {
			err := command.callback(&config, words[1:])
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
	callback    func(config *cliCommandConfig, params []string) error
}

type cliCommandConfig struct {
	Previous *string
	Next     *string
	Client   *pokeapi.Client
	Pokedex  *pokedex.Pokedex
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
		"explore": {
			name:        "explore",
			description: "Displays the pokemon in a given area (name)",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the given pokemon (name)",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays the details of the given pokemon (name)",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all the pokemon in the pokedex",
			callback:    commandPokedex,
		},
	}
}
