package main

import "github.com/madraceee/pokedexcli/internal/pokeapi"

type config struct {
	newURL        *string
	prevURL       *string
	pokeapiClient pokeapi.PokeAPIClient
}

func main() {
	pokeApi := pokeapi.GetNewClient()
	newURL := pokeapi.GetInitialLocationURL()
	prevURL := ""
	cfg := config{
		newURL:        &newURL,
		prevURL:       &prevURL,
		pokeapiClient: pokeApi,
	}
	StartRepl(&cfg)
}
