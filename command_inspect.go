package main

import (
	"errors"
	"fmt"

	"github.com/MaciejSieradz/pokedexcli/internal/pokeapi"
)

func commandInspect(config *config, params ...string) error {
	if len(params) != 1 {
		return errors.New("You must provie a pokemon name")
	}

	name := params[0]

	if val, ok := config.caughtPokemon[name]; ok {
		printInspect(val)
		return nil
	}
	fmt.Println("You have not caught this pokemon")
	return nil
}

func printInspect(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
}
