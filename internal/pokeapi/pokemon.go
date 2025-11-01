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

type PokemonDetails struct {
	Name           string `json:"name"`
	Url            string `json:"url"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
}

type Stat struct {
	Value int
	Name  string
}

func (s *Stat) UnmarshalJSON(data []byte) error {
	var decoded struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	}

	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	s.Name = decoded.Stat.Name
	s.Value = decoded.BaseStat

	return nil
}

type Type struct {
	Slot int
	Name string
}

func (t *Type) UnmarshalJSON(data []byte) error {
	var decoded struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	t.Name = decoded.Type.Name
	t.Slot = decoded.Slot

	return nil
}
