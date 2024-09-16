package main

import (
	"time"

	"github.com/colehuntley83/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeApiClient pokeapi.Client
	nextLocation  *string
	prevLocation  *string
}

func main() {

	client := pokeapi.NewClient(5 * time.Second)
	config := config{
		pokeApiClient: client,
	}
	runRepl(&config)

}
