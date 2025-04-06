package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("you don't have this pokemon: %s", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	fmt.Println("Abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf(" - %s\n", ability.Ability.Name)
	}
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	fmt.Println("Moves:")
	for _, move := range pokemon.Moves {
		fmt.Printf(" - %s\n", move.Move.Name)
	}

	return nil
}
