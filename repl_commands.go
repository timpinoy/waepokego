package main

import (
	"fmt"
	"os"
)

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
	if config.Next == nil && config.Previous != nil {
		fmt.Println("No next map data available")
		return nil
	}
	locs, err := config.Client.GetLocationAreas(config.Next)
	if err != nil {
		fmt.Println("Error when retrieving location-area")
	}
	config.Previous = config.Next
	config.Next = locs.Next
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
	locs, err := config.Client.GetLocationAreas(config.Previous)
	if err != nil {
		fmt.Println("Error when retrieving location-area")
	}
	config.Next = config.Previous
	config.Previous = locs.Previous
	for i := 0; i < len(locs.Results); i++ {
		fmt.Println(locs.Results[i].Name)
	}
	return nil
}
