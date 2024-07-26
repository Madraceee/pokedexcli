package main

import (
	"github.com/madraceee/pokedexcli/internal/pokeapi"
	"time"
)

type config struct {
	newURL        *string
	prevURL       *string
	pokeapiClient pokeapi.PokeAPIClient
	pokemons      map[string]pokeapi.Pokemon
}

func main() {
	pokeApi := pokeapi.GetNewClient(time.Minute * 1)
	newURL := pokeapi.GetInitialLocationURL()
	prevURL := ""
	cfg := config{
		newURL:        &newURL,
		prevURL:       &prevURL,
		pokeapiClient: pokeApi,
		pokemons:      make(map[string]pokeapi.Pokemon),
	}
	StartRepl(&cfg)
}
