package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, parameters ...string) error {
	if cfg.newURL == nil {
		return errors.New("No new area")
	}

	client := cfg.pokeapiClient
	list, err := client.GetLocationList(*cfg.newURL)
	if err != nil {
		return err
	}

	cfg.newURL = &list.Next
	cfg.prevURL = &list.Previous

	fmt.Println()
	for _, v := range list.Results {
		fmt.Println(v.Name)
	}
	fmt.Println()

	return nil
}

func commandMapb(cfg *config, parameters ...string) error {
	if len(*cfg.prevURL) == 0 {
		return errors.New("No previous area")
	}

	client := cfg.pokeapiClient
	list, err := client.GetLocationList(*cfg.prevURL)
	if err != nil {
		return err
	}

	cfg.newURL = &list.Next
	cfg.prevURL = &list.Previous

	fmt.Println()
	for _, v := range list.Results {
		fmt.Println(v.Name)
	}
	fmt.Println()

	return nil
}
