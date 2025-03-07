package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("No parameter provided")
	}
	pokemonsResp, err := config.pokeapiClient.ListPokemons(params[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", params[0])
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonsResp.Pokemons {
		fmt.Printf("   - %v\n", pokemon.PokemonDetails.Name)
	}

	return nil
}
