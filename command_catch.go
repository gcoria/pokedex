package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	randNum := rand.Intn(pokemon.BaseExperience)
	threshold := 60

	if randNum > threshold {
		return fmt.Errorf("Faugh! You didn't catch the pokemon: %s", pokemon.Name)
	}

	fmt.Printf("You caught the pokemon: %s\n", pokemon.Name)

	return nil
}
