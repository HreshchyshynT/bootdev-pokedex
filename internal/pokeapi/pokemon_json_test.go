package pokeapi

import "testing"

const inputJson = `
{
	"pokemon": {
		"name": "tentacool",
		"url": "https://pokeapi.co/api/v2/pokemon/72/"
	},
	"version_details": [
	{
		"version": {
			"name": "diamond",
			"url": "https://pokeapi.co/api/v2/version/12/"
		},
		"max_chance": 60,
		"encounter_details": [
		{
			"min_level": 20,
			"max_level": 30,
			"condition_values": [],
			"chance": 60,
			"method": {
				"name": "surf",
				"url": "https://pokeapi.co/api/v2/encounter-method/5/"
			}
		}
		]
	}
	]
}
`

func TestPokemonFromJson(t *testing.T) {
	input := []byte(inputJson)
	expected := Pokemon{
		Name: "tentacool",
		Url:  "https://pokeapi.co/api/v2/pokemon/72/",
	}

	actual := &Pokemon{}

	err := actual.UnmarshalJSON(input)
	if err != nil {
		t.Errorf("Expected to unmarshal without error: %v\n", err)
	}

	if *actual != expected {
		t.Errorf("Expected %v, received: %v\n", expected, actual)
	}
}
