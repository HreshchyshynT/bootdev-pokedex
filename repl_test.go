package main

import (
	"slices"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{input: "Hello World", output: []string{"hello", "world"}},
		{input: "HELLO WORLD  ", output: []string{"hello", "world"}},
		{input: "hello world", output: []string{"hello", "world"}},
		{input: "   HELLO     world     ", output: []string{"hello", "world"}},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if !slices.Equal(actual, c.output) {
			t.Errorf("expected: %v actual: %v", c.output, actual)
		}
	}
}
