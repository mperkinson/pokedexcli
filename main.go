package main

import "github.com/mperkinson/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient   pokeapi.Client
	prevLocationURL *string
	nextLocationURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
