package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	locationAreasUrl = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
)

type Client struct {
	httpClient *http.Client
	cache      *Cache
}

func NewClient() *Client {
	return &Client{
		httpClient: http.DefaultClient,
		cache:      NewCache(),
	}
}

func (c *Client) String() string {
	return fmt.Sprintf("PokedexApiClient: http is nil %v", c.httpClient == nil)
}

func (c *Client) GetAreas(url string) (Response[LocationArea], error) {
	requestUrl := locationAreasUrl
	if len(url) > 0 {
		requestUrl = url
	}

	var response Response[LocationArea]

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

type Response[T any] struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

func (r Response[T]) String() string {
	return fmt.Sprintf("pokeapi.Response(count=%v, next=%v, previous=%v, results len=%v)\n", r.Count, r.Next, r.Previous, len(r.Results))
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (la LocationArea) String() string {
	return fmt.Sprintf("LocationArea(name=%v, url=%v)", la.Name, la.Url)
}
