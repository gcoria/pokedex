package pokeapi

import (
	"context"
)

func (c *Client) GetPokemon(ctx context.Context, pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullUrl := baseUrl + endpoint

	pokemon := Pokemon{}
	err := c.CachingGetWithContext(ctx, fullUrl, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
