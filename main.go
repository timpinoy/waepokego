package main

import (
	"github.com/timpinoy/waepokego/internal/pokeapi"
	"github.com/timpinoy/waepokego/internal/pokedex"
)

func main() {
	client := pokeapi.NewClient()
	dex := pokedex.New()
	config := NewConfig(client, dex)
	runREPL(config)
}
