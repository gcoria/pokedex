package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func callbackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure resources are cleaned up

	pokemon, err := cfg.pokeapiClient.GetPokemon(ctx, pokemonName)
	if err != nil {
		return err
	}

	randNum := rand.Intn(pokemon.BaseExperience)
	threshold := 60

	if randNum > threshold {
		return fmt.Errorf("Faugh! You didn't catch the pokemon: %s", pokemon.Name)
	}

	cfg.caughtPokemon[pokemonName] = pokemon

	fmt.Printf("You caught the pokemon: %s\n", pokemon.Name)

	return nil
}
