package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	config := NewConfig()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		parts := cleanInput(userInput)
		if len(parts) == 0 {
			continue
		}
		if command, ok := supportedCommands[parts[0]]; ok {
			err := command.callback(parts[1:], config)
			if err != nil {
				fmt.Printf("Error when executing command %v: %v\n", command.name, err)
			}
		} else {
			fmt.Println("Unknown command")
		}
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
