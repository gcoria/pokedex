package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	fmt.Println("Pokedex:")
	fmt.Println("You have caught:")
	for pokemonName := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemonName)
	}

	return nil
}
