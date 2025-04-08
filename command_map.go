package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func callbackMap(cfg *config, args ...string) error {

	// Create context with 15 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure resources are cleaned up

	resp, err := cfg.pokeapiClient.ListLocationAreas(ctx, cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)

	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}

func callbackMapPrevious(cfg *config, args ...string) error {

	if cfg.previousLocationAreaURL == nil {
		return errors.New("You are at the beggining")
	}

	// Create context with 15 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure resources are cleaned up

	resp, err := cfg.pokeapiClient.ListLocationAreas(ctx, cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)

	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}
