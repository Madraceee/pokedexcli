package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, parameters ...string) error {
	pokemons := cfg.pokemons

	if len(pokemons) == 0 {
		return errors.New("No pokemons present in the inventory")
	}

	fmt.Println("Your Pokedex:")
	for _, v := range pokemons {
		fmt.Println("\t- " + v.Name)
	}
	fmt.Println()
	return nil
}
