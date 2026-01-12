package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands(cfg)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanWords := cleanInput(text)

		if len(cleanWords) == 0 {
			continue
		}

		command := cleanWords[0]
		args := []string{}
		if len(cleanWords) > 1 {
			args = cleanWords[1:]
		}

		cmd, ok := commands[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(str string) []string {
	lowerCaseString := strings.ToLower(str)
	words := strings.Fields(lowerCaseString)
	return words
}
