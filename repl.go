package main

import (
	"bufio"
	"fmt"
	"os"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommandMappings()
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		command, present := commands[scanner.Text()]

		if present == false {
			fmt.Println("Wrong command")
		} else {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
