package main

import (
	"fmt"
	"math/rand"
	"os"
)

func commandHelp(config *cliCommandConfig, params []string) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	cliCommands := getCliCommands()
	for _, command := range cliCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Print("\n")
	return nil
}

func commandExit(config *cliCommandConfig, params []string) error {
	os.Exit(0)
	return nil
}

func commandMap(config *cliCommandConfig, params []string) error {
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

func commandMapb(config *cliCommandConfig, params []string) error {
	if config.Previous == nil {
		fmt.Println("No previous map data available")
		return nil
	}
	locs, err := config.Client.GetLocationAreas(config.Previous)
	if err != nil {
		fmt.Println("Error when retrieving location-area")
		return err
	}
	config.Next = config.Previous
	config.Previous = locs.Previous
	for i := 0; i < len(locs.Results); i++ {
		fmt.Println(locs.Results[i].Name)
	}
	return nil
}

func commandExplore(config *cliCommandConfig, params []string) error {
	fmt.Println("Exploring " + params[0] + "...")
	loc, err := config.Client.GetLocationArea(params[0])
	if err != nil {
		fmt.Println("Error when retrieving information")
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range loc.PokemonEncounters {
		fmt.Printf("\t- %s\n", pokemonEncounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(config *cliCommandConfig, params []string) error {
	fmt.Println("Throwing a Pokeball at " + params[0] + "...")
	pokemon, err := config.Client.GetPokemon(params[0])
	if err != nil {
		fmt.Println("Error when retrieving information")
		return err
	}
	rng := rand.Intn(400 + pokemon.BaseExperience)
	if rng < 200 {
		fmt.Printf("%s escaped\n", params[0])
	} else {
		fmt.Printf("%s was caught!\n", params[0])
		fmt.Printf("You may now inspect it with the inspect command.\n")
		config.Pokedex.Add(pokemon)
	}
	return nil
}

func commandInspect(config *cliCommandConfig, params []string) error {
	pokemon, err := config.Pokedex.Get(params[0])
	if err != nil {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typ := range pokemon.Types {
		fmt.Printf("\t- %s\n", typ.Type.Name)
	}
	return nil
}

func commandPokedex(config *cliCommandConfig, params []string) error {
	fmt.Printf("Your Pokedex:\n")
	for _, pokemon := range config.Pokedex.List() {
		fmt.Printf("\t- %s\n", pokemon.Name)
	}
	return nil
}
