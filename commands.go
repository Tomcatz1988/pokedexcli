package main

import (
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
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations in the Pokemon world",
			callback:    commandMapBack,
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
	batch, err := pokeapi.GetLocationBatch(conf.Next)
	if err != nil {
		return fmt.Errorf("commandMap(): %w", err)
	}

	for _, location := range(batch.Results) {
		fmt.Println(location.Name)
	}
	if batch.Next != nil {
		conf.Next = *batch.Next
	}
	if batch.Previous != nil {
		conf.Previous = *batch.Previous
	}
	return nil
}

func commandMapBack(conf *config) error {
	batch, err := pokeapi.GetLocationBatch(conf.Previous)
	if err != nil {
		return fmt.Errorf("commandMapBack(): %w: ",err)
	}

	for _, location := range(batch.Results) {
		fmt.Println(location.Name)
	}
	if batch.Next != nil {
		conf.Next = *batch.Next
	}
	if batch.Previous != nil {
		conf.Previous = *batch.Previous
	}
	return nil
}
