package main

import (
	"errors"
	"fmt"
	"os"
	pokeapi "github.com/Tomcatz1988/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

func getCommandsRegistry() (reg map[string]cliCommand) {
	reg = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations in the Pokemon world",
			callback:    commandMap,
		},
	}
	return reg
}

func commandExit(conf *config) error {
	_ = conf
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	_ = conf
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	reg := getCommandsRegistry()
	for _, command := range reg {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(conf *config) error {
	_ = conf
	_, _ = pokeapi.GetLocationBatch(conf.Next)
	return errors.New("command not implemented yet")
}
