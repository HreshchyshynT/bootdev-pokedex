package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hreshchyshynt/pokedex/internal/pokecache"
)

const (
	locationAreasUrl       = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	locationAreaDetailsUrl = "https://pokeapi.co/api/v2/location-area/"
)

type Client struct {
	httpClient *http.Client
	cache      pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		httpClient: http.DefaultClient,
		cache:      pokecache.NewCache(time.Duration(10) * time.Second),
	}
}

func (c *Client) String() string {
	return fmt.Sprintf("PokedexApiClient: http is nil %v", c.httpClient == nil)
}

func (c *Client) GetAreas(url string) (Response[LocationAreaShort], error) {
	requestUrl := locationAreasUrl
	if len(url) > 0 {
		requestUrl = url
	}

	var response Response[LocationAreaShort]

	if data, ok := c.cache.Get(requestUrl); ok {
		err := json.Unmarshal(data, &response)
		if err == nil {
			return response, nil
		}
	}
	res, err := c.httpClient.Get(requestUrl)
	if err != nil {
		return response, fmt.Errorf("can not make request for location areas: %v", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return response, fmt.Errorf("error reading reading response: %v\n", err)
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, fmt.Errorf("error unmarshaling response: %v\n", err)
	}

	c.cache.Put(requestUrl, data)

	return response, nil
}

func (c *Client) GetPokemonsOnLocation(name string) (LocationAreaDetails, error) {
	requestUrl := locationAreaDetailsUrl + name

	var response LocationAreaDetails

	if data, ok := c.cache.Get(requestUrl); ok {
		err := json.Unmarshal(data, &response)
		if err == nil {
			return response, nil
		}
	}
	res, err := c.httpClient.Get(requestUrl)
	if err != nil {
		return response, fmt.Errorf("can not make request for location area details: %v", err)
	}

	if res.StatusCode == http.StatusNotFound {
		return response, fmt.Errorf("Location \"%v\" not found", name)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return response, fmt.Errorf("error reading reading response: %v\n", err)
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, fmt.Errorf("error unmarshaling response: %v\n", err)
	}

	c.cache.Put(requestUrl, data)

	return response, nil
}
