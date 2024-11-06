package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		if command, ok := getCliCommands()[userInput]; ok {
			err := command.callback()
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
	callback    func() error
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
	}
}

func commandHelp() error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	cliCommands := getCliCommands()
	for _, command := range cliCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Print("\n")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
