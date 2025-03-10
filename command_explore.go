package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a location area name")
	}

	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
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
