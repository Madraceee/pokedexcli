package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommandMappings() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit pokedex",
			callback:    commandExit,
		},
	}
}
