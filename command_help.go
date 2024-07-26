package main

import "fmt"

func commandHelp(cfg *config, parameters ...string) error {
	commands := getCommandMappings()

	fmt.Println()
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	fmt.Println()
	return nil
}
