package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(location string) (RespPokemonLocations, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		pokemonResponse := RespPokemonLocations{}
		err := json.Unmarshal(val, &pokemonResponse)
		if err != nil {
			return RespPokemonLocations{}, err
		}

		return pokemonResponse, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonLocations{}, err
	}

	pokemonsResponse := RespPokemonLocations{}
	err = json.Unmarshal(data, &pokemonsResponse)
	if err != nil {
		return RespPokemonLocations{}, err
	}

	return pokemonsResponse, nil
}
