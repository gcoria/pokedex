package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func callbackExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a location area name")
	}

	locationAreaName := args[0]

	// Create context with 15 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure resources are cleaned up

	locationArea, err := cfg.pokeapiClient.GetLocationArea(ctx, locationAreaName)
	if err != nil {
		return err
	}

	fmt.Println("Location Area:")
	fmt.Printf(" - %s\n", locationArea.Name)

	fmt.Println("Pokemon:")
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
