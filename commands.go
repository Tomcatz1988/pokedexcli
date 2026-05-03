package main

import (
	"errors"
	"fmt"
	"os"

	pokeapi "internal/pokeapi"
	pokecache "internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config, cache *pokecache.Cache, args []string) error
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
		"explore": {
			name:        "explore",
			description: "Displays all species of pokemone in the specified area",
			callback:    commandExplore,
		},
	}
	return reg
}

func commandExit(conf *config, cache *pokecache.Cache, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config, cache *pokecache.Cache, args []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	reg := getCommandsRegistry()
	for _, command := range reg {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(conf *config, cache *pokecache.Cache, args []string) error {
	batch, err := pokeapi.GetAreaBatch(conf.Next, cache)
	if err != nil {
		return fmt.Errorf("commandMap: %w", err)
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

func commandMapBack(conf *config, cache *pokecache.Cache, args []string) error {
	batch, err := pokeapi.GetAreaBatch(conf.Previous, cache)
	if err != nil {
		return fmt.Errorf("commandMapBack: %w: ",err)
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


func commandExplore(conf *config, cache *pokecache.Cache, args []string) error {
	if len(args) == 0 {
		return errors.New("'explore' requires a [location] as an argument - explore [location]")
	}

	fmt.Printf("Exploring %v...\n", args[0])
	targetURL := locationURL + args[0] + "/"
	info, err := pokeapi.GetAreaInfo(targetURL, cache)
	if err != nil {
		return fmt.Errorf("commandExplore: %w", err)
	}

	encounters := info.PokemonEncounters
	for _, encounter := range(encounters) {
		pokemon := encounter.Pokemon.Name
		fmt.Printf("- %v\n", pokemon)
	}
	return nil
}
