package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/timpinoy/waepokego/internal/pokeapi"
)

const basePokeapiUrl string = "https://pokeapi.co/api/v2/"

func main() {
	areaLocationURL := basePokeapiUrl + "location-area"
	mapCommandConfig := cliCommandConfig{Next: &areaLocationURL}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		if command, ok := getCliCommands()[userInput]; ok {
			err := command.callback(&mapCommandConfig)
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

func commandHelp(config *cliCommandConfig) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	cliCommands := getCliCommands()
	for _, command := range cliCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Print("\n")
	return nil
}

func commandExit(config *cliCommandConfig) error {
	os.Exit(0)
	return nil
}

func commandMap(config *cliCommandConfig) error {
	if config.Next == nil {
		fmt.Println("No next map data available")
		return nil
	}
	locs, err := pokeapi.GetLocationAreas(*config.Next)
	if err != nil {
		fmt.Println("Error when retrieving location-area")
	}
	config.Next = locs.Next
	config.Previous = locs.Previous
	for i := 0; i < len(locs.Results); i++ {
		fmt.Println(locs.Results[i].Name)
	}
	return nil
}

func commandMapb(config *cliCommandConfig) error {
	if config.Previous == nil {
		fmt.Println("No previous map data available")
		return nil
	}
	locs, err := pokeapi.GetLocationAreas(*config.Previous)
	if err != nil {
		fmt.Println("Error when retrieving location-area")
	}
	config.Next = locs.Next
	config.Previous = locs.Previous
	for i := 0; i < len(locs.Results); i++ {
		fmt.Println(locs.Results[i].Name)
	}
	return nil
}
