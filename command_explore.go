package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, parameters ...string) error {
	if len(parameters) == 0 {
		return errors.New("Enter Location name")
	}

	client := cfg.pokeapiClient
	location, err := client.ExploreLocation(parameters[0])
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + location.Name + "...")
	fmt.Println("Found Pokemon:")

	for _, v := range location.PokemonEncounters {
		fmt.Println("- " + v.Pokemon.Name)
	}

	return nil
}
