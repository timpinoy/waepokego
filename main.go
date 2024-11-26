package main

import "github.com/timpinoy/waepokego/internal/pokeapi"

func main() {
	client := pokeapi.NewClient()
	config := NewConfig(client)
	runREPL(config)
}
