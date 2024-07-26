package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommandMappings()
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		command, present := commands[input[0]]

		if present == false {
			fmt.Println("Wrong command")
		} else {
			err := command.callback(cfg, input[1:]...)
			if err != nil {
				fmt.Println()
				fmt.Println(err)
				fmt.Println()
			}
		}
	}
}
