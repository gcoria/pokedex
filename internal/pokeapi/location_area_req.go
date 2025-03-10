package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	if pageURL != nil {
		fullUrl = *pageURL
	}

	if dat, ok := c.cache.Get(fullUrl); ok {
		locatationAreaResp := LocationAreaResp{}
		err := json.Unmarshal(dat, &locatationAreaResp)
		if err != nil {
			return LocationAreaResp{}, err
		}

		return locatationAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code[%v]", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locatationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locatationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullUrl, data)

	return locatationAreaResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullUrl := baseUrl + endpoint

	if dat, ok := c.cache.Get(fullUrl); ok {
		locatationAreaResp := LocationArea{}
		err := json.Unmarshal(dat, &locatationAreaResp)
		if err != nil {
			return LocationArea{}, err
		}

		return locatationAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code[%v]", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locatationAreaResp := LocationArea{}
	err = json.Unmarshal(data, &locatationAreaResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullUrl, data)

	return locatationAreaResp, nil
}
