package pokeapi

import "fmt"

type Response[T any] struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

func (r Response[T]) String() string {
	return fmt.Sprintf("pokeapi.Response(count=%v, next=%v, previous=%v, results len=%v)\n", r.Count, r.Next, r.Previous, len(r.Results))
}
