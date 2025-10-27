package main

import (
	"fmt"
	"github.com/hreshchyshynt/pokedex/internal/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func printHelp(config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, v := range supportedCommands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}

func displayMap(config *Config) error {
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

func displayMapb(config *Config) error {
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
