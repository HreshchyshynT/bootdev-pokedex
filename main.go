package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		parts := cleanInput(userInput)
		fmt.Printf("Your command was: %v\n", parts[0])
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
