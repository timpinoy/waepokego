package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/timpinoy/waepokego/internal/pokecache"
	"io"
	"log"
	"net/http"
	"time"
)

const baseURL string = "https://pokeapi.co/api/v2"

type Client struct {
	cache *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		cache: pokecache.NewCache(time.Second * 30),
	}
}

func (c *Client) GetLocationAreas(url *string) (PokeAPILocationAreas, error) {
	locationAreas := PokeAPILocationAreas{}
	queryURL := baseURL + "/location-area"
	if url != nil {
		queryURL = *url
	}

	data, ok := c.cache.Get(queryURL)
	if ok {
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			fmt.Println(err)
		}
		return locationAreas, nil
	}

	res, err := http.Get(queryURL)
	if err != nil {
		return PokeAPILocationAreas{}, fmt.Errorf("Error during GET: %v", err)
	}
	data, err = io.ReadAll(res.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}
	if err != nil {
		return PokeAPILocationAreas{}, fmt.Errorf("Error during response read: %v", err)
	}

	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		fmt.Println(err)
	}

	c.cache.Add(queryURL, data)

	return locationAreas, nil
}

func (c *Client) GetLocationArea(param string) (PokeAPILocationArea, error) {
	locationArea := PokeAPILocationArea{}
	queryURL := baseURL + "/location-area/" + param

	data, ok := c.cache.Get(queryURL)
	if ok {
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			fmt.Println(err)
		}
		return locationArea, nil
	}

	res, err := http.Get(queryURL)
	if err != nil {
		return PokeAPILocationArea{}, fmt.Errorf("Error during GET: %v", err)
	}
	data, err = io.ReadAll(res.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}
	if err != nil {
		return PokeAPILocationArea{}, fmt.Errorf("Error during response read: %v", err)
	}

	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		fmt.Println(err)
	}

	c.cache.Add(queryURL, data)

	return locationArea, nil
}

func (c *Client) GetPokemon(param string) (Pokemon, error) {
	pokemon := Pokemon{}
	queryURL := baseURL + "/pokemon/" + param

	data, ok := c.cache.Get(queryURL)
	if ok {
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			fmt.Println(err)
		}
		return pokemon, nil
	}

	res, err := http.Get(queryURL)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error during GET: %v", err)
	}
	data, err = io.ReadAll(res.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error during response read: %v", err)
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		fmt.Println(err)
	}

	c.cache.Add(queryURL, data)

	return pokemon, nil
}
