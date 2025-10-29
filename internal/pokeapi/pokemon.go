package pokeapi

import (
	"encoding/json"
)

type Pokemon struct {
	Name string
	Url  string
}

func (p *Pokemon) UnmarshalJSON(data []byte) error {
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
