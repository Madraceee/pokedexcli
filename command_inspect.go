package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, parameters ...string) error {
	if len(parameters) == 0 {
		return errors.New("Enter pokemon name")
	}

	if pokemon, present := cfg.pokemons[parameters[0]]; present == false {
		return errors.New("Pokemon not present in the inventory")
	} else {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, v := range pokemon.Stats {
			fmt.Printf("\t-%s: %d\n", v.Stat.Name, v.BaseStat)
		}
		fmt.Println("Types:")
		for _, v := range pokemon.Types {
			fmt.Println("\t- " + v.Type.Name)
		}
	}

	return nil
}
