package pokeapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gcoria/pokedex/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c Client) Get(url string, target interface{}) error {
	// Create a timeout context for the default Get method
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.GetWithContext(ctx, url, target)
}

func (c Client) GetWithContext(ctx context.Context, url string, target interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(target)
}

// CachingGetWithContext performs a GET request with caching support.
// It handles checking the cache first and updating it after a successful request.
func (c *Client) CachingGetWithContext(ctx context.Context, url string, target interface{}) error {
	// Check cache first
	if dat, ok := c.cache.Get(url); ok {
		return json.Unmarshal(dat, target)
	}

	// Make the request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return fmt.Errorf("bad status code[%v]", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse JSON
	err = json.Unmarshal(data, target)
	if err != nil {
		return err
	}

	// Add to cache
	c.cache.Add(url, data)

	return nil
}
