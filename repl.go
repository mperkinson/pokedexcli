package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerCaseString := strings.ToLower(text)
	words := strings.Fields(lowerCaseString)
	return words
}
