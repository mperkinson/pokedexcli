package main

import (
	"time"

	"github.com/mperkinson/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	prevLocationURL *string
	nextLocationURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startRepl(&cfg)
}
