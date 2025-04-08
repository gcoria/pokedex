package pokeapi

import (
	"context"
)

func (c *Client) ListLocationAreas(ctx context.Context, pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	if pageURL != nil {
		fullUrl = *pageURL
	}

	locationAreaResp := LocationAreaResp{}
	err := c.CachingGetWithContext(ctx, fullUrl, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return locationAreaResp, nil
}

func (c *Client) GetLocationArea(ctx context.Context, locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullUrl := baseUrl + endpoint

	locationArea := LocationArea{}
	err := c.CachingGetWithContext(ctx, fullUrl, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}
