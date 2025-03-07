package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(config *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("You must provide a pokemon name")
	}
	pokemon, err := config.pokeapiClient.Pokemon(params[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	random := rand.IntN(pokemon.BaseExperience)
	if random > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	config.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
