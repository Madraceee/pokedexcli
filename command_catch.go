package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, parameters ...string) error {
	if len(parameters) == 0 {
		return errors.New("Enter pokemon name")
	}

	if _, present := cfg.pokemons[parameters[0]]; present == true {
		return errors.New("Pokemon already caught!")
	}

	client := cfg.pokeapiClient
	pokemon, err := client.PokemonDetails(parameters[0])
	if err != nil {
		return err
	}

	randNumber := rand.Intn(pokemon.BaseExperience)

	fmt.Println()
	fmt.Println("Throwing a Pokeball at pikachu...")
	if randNumber >= 50 {
		fmt.Println(pokemon.Name + " was caught!")
		cfg.pokemons[parameters[0]] = pokemon
	} else {
		fmt.Println(pokemon.Name + " escaped!")
	}
	fmt.Println()

	return nil
}
