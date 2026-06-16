package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonData, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if cachedData, ok := c.cache.Get(url); ok{
		var pokemonData PokemonData
		if err := json.Unmarshal(cachedData, &pokemonData); err != nil {
			return PokemonData{},fmt.Errorf("Error: %v", err)
		}

		return pokemonData, nil
	}

	// make the GET request using c.httpClient
	res, err := c.httpClient.Get(url)
	if err != nil {
		return PokemonData{}, fmt.Errorf("Error: %v", err)
	}
	defer res.Body.Close()
	
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonData{},fmt.Errorf("Error: %v", err)
	}

	c.cache.Add(url, bodyBytes)

	// 4. unmarshal into a location
	var pokemonData PokemonData
	if err := json.Unmarshal(bodyBytes, &pokemonData); err != nil {
		return PokemonData{},fmt.Errorf("Error: %v", err)
	}

	return pokemonData, nil
}
