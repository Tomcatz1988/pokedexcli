package main

import (
	"errors"
	"fmt"
	"os"

	pokeapi "internal/pokeapi"
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
		"explore": {
			name:        "explore",
			description: "Displays all species of pokemon in the specified area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to add the specified pokemon to the users collection",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "View information about the specified pokemon if it exists in the pokedex",
			callback:    commandInspect,
		},
	}
	return reg
}


func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp(conf *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	reg := getCommandsRegistry()
	for _, command := range reg {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}


func commandMap(conf *config) error {
	cache := conf.cache
	batch, err := pokeapi.GetAreaBatch(conf.Next, cache)
	if err != nil {
		return fmt.Errorf("error: commandMap: %w", err)
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
	cache := conf.cache
	batch, err := pokeapi.GetAreaBatch(conf.Previous, cache)
	if err != nil {
		return fmt.Errorf("error: commandMapBack: %w: ",err)
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


func commandExplore(conf *config) error {
	cache := conf.cache
	args := conf.args
	if len(args) == 0 {
		return errors.New("'explore' requires a [location] as an argument - explore [location]")
	}

	fmt.Printf("Exploring %v...\n", args[0])
	targetURL := baseURL + locationURL + args[0] + "/"
	info, err := pokeapi.GetAreaInfo(targetURL, cache)
	if err != nil {
		return fmt.Errorf("error: commandExplore: %w", err)
	}

	encounters := info.PokemonEncounters
	for _, encounter := range(encounters) {
		pokemon := encounter.Pokemon.Name
		fmt.Printf("- %v\n", pokemon)
	}
	return nil
}


func commandCatch(conf *config) error {
	cache := conf.cache
	args := conf.args
	pokedex := conf.pokedex
	if len(args) == 0 {
		return errors.New("'catch' requires a [pokemon] as an argument - explore [pokemon]")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", name)
	targetURL := baseURL + pokemonURL + name + "/"
	info, err := pokeapi.GetPokemon(targetURL, cache)
	if err != nil {
		return fmt.Errorf("error: commandCatch: %w", err)
	}

	if catch(info, pokedex) {
		fmt.Printf("%v was caught!\n", name)
	} else {
		fmt.Printf("%v escaped!\n", name)
	}
	return nil
}


func commandInspect(conf *config) error {
	args := conf.args
	pokedex := conf.pokedex
	if len(args) == 0 {
		return errors.New("'inspect' requires a [pokemon] as an argument - explore [pokemon]")
	}

	name := conf.args[0]
	pokemon, exists := pokedex[name]
	if !exists {
		return fmt.Errorf("%s has not been caught yet", name)
	}

	fmt.Printf("Name: %s\nHeight: %v\nWeight: %v\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range(pokemon.Stats) {
		fmt.Printf("  -%s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range(pokemon.Types) {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
