package main

import (
	"time"

	"github.com/mperkinson/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	prevLocationURL *string
	nextLocationURL *string
	pokedex         map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
