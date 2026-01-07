package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Empty String",
			input:    "",
			expected: []string{},
		},
		{
			name:     "Space Only",
			input:    " ",
			expected: []string{},
		},
		{
			name:     "Two Words",
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "Extra Spaces",
			input:    " hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "Mixed Case",
			input:    "HeLlO WORld",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cleanInput(c.input)

			if len(actual) != len(c.expected) {
				t.Fatalf(
					"cleanInput(%q) length mismatch\nexpected: %d (%v)\nactual:   %d (%v)",
					c.input,
					len(c.expected),
					c.expected,
					len(actual),
					actual,
				)
			}

			for i := range c.expected {
				if actual[i] != c.expected[i] {
					t.Fatalf(
						"cleanInput(%q) mismatch at index %d\nexpected: %q\nactual:   %q\nfull expected: %v\nfull actual:   %v",
						c.input,
						i,
						c.expected[i],
						actual[i],
						c.expected,
						actual,
					)
				}
			}
		})
	}
}
