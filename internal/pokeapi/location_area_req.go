package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

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

	return locatationAreaResp, nil
}
