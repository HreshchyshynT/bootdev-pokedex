package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
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
