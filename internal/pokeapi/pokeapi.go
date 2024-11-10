package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PokeAPILocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (PokeAPILocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return PokeAPILocationAreas{}, fmt.Errorf("Error during GET: %v", err)
	}
	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}
	if err != nil {
		return PokeAPILocationAreas{}, fmt.Errorf("Error during response read: %v", err)
	}
	locationAreas := PokeAPILocationAreas{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		fmt.Println(err)
	}
	return locationAreas, nil
}
