package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/hreshchyshynt/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	// command payload
	callback func(args []string, config *Config) error
}

type Config struct {
	Next string
	Prev string
}

func NewConfig() *Config {
	return &Config{}
}

var supportedCommands map[string]cliCommand
var apiClient *pokeapi.Client
var pokedex map[string]pokeapi.PokemonDetails

// run automatically before main function
// we have to init supportedCommands here to avoid
// initialization cycle
func init() {
	apiClient = pokeapi.NewClient()
	pokedex = make(map[string]pokeapi.PokemonDetails)
	supportedCommands = make(map[string]cliCommand)

	supportedCommands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	supportedCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    printHelp,
	}
	supportedCommands["map"] = cliCommand{
		name:        "map",
		description: "Displays the names of 20 location areas in the Pokemon world",
		callback:    displayMap,
	}
	supportedCommands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the names of previous 20 location areas in the Pokemon world",
		callback:    displayMapb,
	}
	supportedCommands["explore"] = cliCommand{
		name:        "explore",
		description: "List of all the Pokémon on the given location",
		callback:    explore,
	}
	supportedCommands["catch"] = cliCommand{
		name:        "catch",
		description: "Catching Pokemon adds them to the user's Pokedex",
		callback:    catchPokemon,
	}
	supportedCommands["inspect"] = cliCommand{
		name:        "inspect",
		description: "Allow players to see details about a Pokemon if they have seen it before (or in our case, caught it)",
		callback:    inspectPokemon,
	}
}

func commandExit(args []string, config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func printHelp(args []string, config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, v := range supportedCommands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}

func displayMap(args []string, config *Config) error {
	res, err := apiClient.GetAreas(config.Next)
	if err != nil {
		return err
	}
	for _, la := range res.Results {
		fmt.Printf("%v\n", la.Name)
	}
	config.Next = res.Next
	config.Prev = res.Previous
	return nil
}

func displayMapb(args []string, config *Config) error {
	res, err := apiClient.GetAreas(config.Prev)
	if err != nil {
		return err
	}
	for _, la := range res.Results {
		fmt.Printf("%v\n", la.Name)
	}
	config.Next = res.Next
	config.Prev = res.Previous
	return nil
}

func explore(args []string, config *Config) error {
	if len(args) == 0 {
		return fmt.Errorf("Location name must be provided.")
	}
	locationName := args[0]
	fmt.Printf("Exploring %v...\n", locationName)

	res, err := apiClient.GetPokemonsOnLocation(locationName)
	if err != nil {
		return err
	}

	if len(res.Pokemons) == 0 {
		fmt.Println("No Pokemons found")
		return nil
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range res.Pokemons {
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return nil
}

func catchPokemon(args []string, config *Config) error {
	if len(args) == 0 {
		return fmt.Errorf("Pokemon name must be provided.")
	}
	name := args[0]

	if _, ok := pokedex[name]; ok {
		fmt.Printf("%v already in pokedex!\n", name)
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	pokemon, err := apiClient.GetPokemonDetails(name)
	if err != nil {
		return err
	}

	var percentRequired float32
	if pokemon.BaseExperience < 50 {
		percentRequired = 0.3
	} else if pokemon.BaseExperience < 90 {
		percentRequired = 0.5
	} else {
		percentRequired = 0.75
	}

	barrier := int(float32(pokemon.BaseExperience) * percentRequired)
	roll := rand.Intn(pokemon.BaseExperience)

	if roll >= barrier {
		pokedex[name] = pokemon
		fmt.Printf("%v was caught!\n", name)
	} else {
		fmt.Printf("%v escaped!\n", name)
	}

	return nil
}

func inspectPokemon(args []string, config *Config) error {
	if len(args) == 0 {
		return fmt.Errorf("Pokemon name must be provided.")
	}
	name := args[0]
	pokemon, ok := pokedex[name]

	if !ok {
		return fmt.Errorf("You didn't see %v yet!", name)
	}

	buffer := &strings.Builder{}

	fmt.Fprintf(buffer, "Name: %v\n", name)
	fmt.Fprintf(buffer, "Height: %v\n", pokemon.Height)
	fmt.Fprintf(buffer, "Weight: %v\n", pokemon.Weight)
	buffer.WriteString("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Fprintf(buffer, "  -%v: %v\n", stat.Name, stat.Value)
	}

	buffer.WriteString("Types:\n")

	for _, t := range pokemon.Types {
		fmt.Fprintf(buffer, "  - %v\n", t.Name)
	}

	fmt.Println(buffer.String())

	return nil
}
