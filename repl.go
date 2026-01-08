package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanWords := cleanInput(text)
		fmt.Printf("Your command was: %s\n", cleanWords[0])
	}
}

func cleanInput(str string) []string {
	lowerCaseString := strings.ToLower(str)
	words := strings.Fields(lowerCaseString)
	return words
}
