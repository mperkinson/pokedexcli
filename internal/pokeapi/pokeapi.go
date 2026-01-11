// Package pokeapi to interact with the Pokeapi API
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mperkinson/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	cache      *pokecache.Cache
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

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache hit!")
		locationAreaResp := LocationAreaResp{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationAreaResp, nil

	}

	fmt.Println("Cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			fmt.Println("Error closing body response", err)
		}
	}()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreaResp, nil
}
