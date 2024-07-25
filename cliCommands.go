package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
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
		"map": {
			name:        "map",
			description: "View next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "View previous 20 locations",
			callback:    commandMapb,
		},
	}
}
