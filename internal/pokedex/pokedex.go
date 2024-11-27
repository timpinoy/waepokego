package pokedex

import (
	"errors"
	"github.com/timpinoy/waepokego/internal/pokeapi"
	"maps"
)

type Pokedex struct {
	list map[string]pokeapi.Pokemon
}

func New() *Pokedex {
	return &Pokedex{
		list: make(map[string]pokeapi.Pokemon),
	}
}

func (p *Pokedex) Get(name string) (pokeapi.Pokemon, error) {
	if pokemon, ok := p.list[name]; ok {
		return pokemon, nil
	} else {
		return pokeapi.Pokemon{}, errors.New("not found")
	}
}

func (p *Pokedex) Add(pokemon pokeapi.Pokemon) {
	p.list[pokemon.Name] = pokemon
}

func (p *Pokedex) Remove(name string) {
	delete(p.list, name)
}

func (p *Pokedex) List() []pokeapi.Pokemon {
	list := make([]pokeapi.Pokemon, len(p.list))
	i := 0
	for v := range maps.Values(p.list) {
		list[i] = v
		i++
	}
	return list
}
