package main

import (
	"fmt"
	"github.com/hreshchyshynt/pokedex/internal/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	// command payload
	callback func([]string, *Config) error
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

// run automatically before main function
// we have to init supportedCommands here to avoid
// initialization cycle
func init() {
	apiClient = pokeapi.NewClient()
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
}

func commandExit(input []string, config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func printHelp(input []string, config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, v := range supportedCommands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}

func displayMap(input []string, config *Config) error {
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

func displayMapb(input []string, config *Config) error {
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

func explore(input []string, config *Config) error {
	if len(input) == 0 {
		return fmt.Errorf("Location name should be provided.")
	}
	locationName := input[0]
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
