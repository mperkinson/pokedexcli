package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanWords := cleanInput(text)

		if len(cleanWords) == 0 {
			continue
		}

		command := cleanWords[0]

		cmd, ok := commands[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback()
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
