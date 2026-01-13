package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name          string
	description   string
	callback      func(*config, ...string) error
	configuration *config
}

func getCommands(cfg *config) map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:          "exit",
			description:   "Exit the Pokedex",
			callback:      commandExit,
			configuration: cfg,
		},
		"help": {
			name:          "help",
			description:   "Displays a help message",
			callback:      commandHelp,
			configuration: cfg,
		},
		"map": {
			name:          "map",
			description:   "Displays next page of location areas",
			callback:      commandMap,
			configuration: cfg,
		},
		"mapb": {
			name:          "mapb",
			description:   "Displays previous page of location areas",
			callback:      commandMapb,
			configuration: cfg,
		},
		"explore": {
			name:          "explore {location area}",
			description:   "Displays Pokemon found in a location area",
			callback:      commandExplore,
			configuration: cfg,
		},
		"catch": {
			name:          "catch {Pokemon name}",
			description:   "Attempt to catch a Pokemon and add it to the Pokedex",
			callback:      commandCatch,
			configuration: cfg,
		},
	}
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for _, c := range getCommands(cfg) {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	cfg.prevLocationURL = resp.Previous
	cfg.nextLocationURL = resp.Next

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("no previous page to return to")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	cfg.prevLocationURL = resp.Previous
	cfg.nextLocationURL = resp.Next

	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	locationAreaName := args[0]

	resp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", resp.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no Pokemon name provided")
	}

	pokemonName := args[0]

	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	const threshold = 50
	randomNum := rand.Intn(resp.BaseExperience)
	if randomNum > threshold {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	cfg.pokedex[pokemonName] = resp
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}
