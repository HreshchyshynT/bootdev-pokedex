package pokeapi

import "fmt"

type LocationAreaShort struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (las LocationAreaShort) String() string {
	return fmt.Sprintf("LocationArea(name=%v, url=%v)", las.Name, las.Url)
}

type LocationAreaDetails struct {
	Id       int            `json:"id"`
	Name     string         `json:"name"`
	Pokemons []PokemonShort `json:"pokemon_encounters"`

	// {
	//   "id": 1,
	//   "name": "canalave-city-area",
	//   "game_index": 1,
	//   "encounter_method_rates": [
	//     {
	//       "encounter_method": {
	//         "name": "old-rod",
	//         "url": "https://pokeapi.co/api/v2/encounter-method/2/"
	//       },
	//       "version_details": [
	//         {
	//           "rate": 25,
	//           "version": {
	//             "name": "platinum",
	//             "url": "https://pokeapi.co/api/v2/version/14/"
	//           }
	//         }
	//       ]
	//     }
	//   ],
	//   "location": {
	//     "name": "canalave-city",
	//     "url": "https://pokeapi.co/api/v2/location/1/"
	//   },
	//   "names": [
	//     {
	//       "name": "",
	//       "language": {
	//         "name": "en",
	//         "url": "https://pokeapi.co/api/v2/language/9/"
	//       }
	//     }
	//   ],
	//   "pokemon_encounters": [
	//     {
	//       "pokemon": {
	//         "name": "tentacool",
	//         "url": "https://pokeapi.co/api/v2/pokemon/72/"
	//       },
	//       "version_details": [
	//         {
	//           "version": {
	//             "name": "diamond",
	//             "url": "https://pokeapi.co/api/v2/version/12/"
	//           },
	//           "max_chance": 60,
	//           "encounter_details": [
	//             {
	//               "min_level": 20,
	//               "max_level": 30,
	//               "condition_values": [],
	//               "chance": 60,
	//               "method": {
	//                 "name": "surf",
	//                 "url": "https://pokeapi.co/api/v2/encounter-method/5/"
	//               }
	//             }
	//           ]
	//         }
	//       ]
	//     }
	//   ]
	// }

}
