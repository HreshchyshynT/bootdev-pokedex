package main

import (
	"bufio"
	"fmt"
	"github.com/hreshchyshynt/pokedex/pokeapi"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	client := pokeapi.NewClient()
	fmt.Printf("pokeapi client: %v\n", client)
	areasResponse, err := client.GetAreas()
	if err != nil {
		fmt.Printf("error getting areas: %v\n", err)
	}

	fmt.Printf("Areas received: %v\n", areasResponse)
	for i, la := range areasResponse.Results {
		fmt.Printf("Item %v: %v\n", i, la)
		if i > 9 {
			break
		}
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		parts := cleanInput(userInput)
		if len(parts) == 0 {
			continue
		}
		if command, ok := supportedCommands[parts[0]]; ok {
			err := command.callback()
			if err != nil {
				fmt.Printf("Error when executing command %v: %v\n", command.name, err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error during reading input: %v\n", err)
	}
}

func cleanInput(text string) []string {
	var result []string

	var builder strings.Builder

	for i, c := range text {
		if c != ' ' {
			builder.WriteRune(c)
		}
		if (c == ' ' || i == len(text)-1) && builder.Len() > 0 {
			result = append(result, strings.ToLower(builder.String()))
			builder.Reset()
		}
	}
	return result
}
