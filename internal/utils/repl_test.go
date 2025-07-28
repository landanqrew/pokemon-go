package utils

import (
	"slices"
	"testing"
)


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "  hello-world ",
			expected: []string{"hello-world"},
		},
		{
			input: "  this is an   input example",
			expected: []string{"this", "is", "an", "input", "example"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if !slices.Equal(actual, c.expected) {
			t.Errorf("cleanInput(%v) = %v; expects %v", c.input, actual, c.expected)
		}
	}
}