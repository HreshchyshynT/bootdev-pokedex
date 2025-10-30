package pokeapi

import (
	"encoding/json"
)

type PokemonShort struct {
	Name string
	Url  string
}

func (p *PokemonShort) UnmarshalJSON(data []byte) error {
	var decoded struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	}

	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	p.Name = decoded.Pokemon.Name
	p.Url = decoded.Pokemon.Url

	return nil
}
