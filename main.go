package main

import (
	"bufio"
	"fmt"
	"os"

	pokeapi "internal/pokeapi"
	pokecache "internal/pokecache"
)

type config struct {
	Next string
	Previous string
	reg map[string]cliCommand
	cache *pokecache.Cache
	args []string
	pokedex map[string]pokeapi.Pokemon
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	reg := getCommandsRegistry()
	cache := pokecache.NewCache(cacheDuration)
	pokedex := make(map[string]pokeapi.Pokemon)
	conf := config{
		Next: baseURL + locationURL,
		Previous: baseURL + locationURL,
		reg: reg,
		cache: &cache,
		pokedex: pokedex,
	}

	for {
		fmt.Print("\nPokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		fmt.Println("")
		if len(words) == 0 {
			continue
		}
		command, exists := reg[words[0]]
		if exists {
			var args []string
			if len(words) > 1 {
				args = words[1:]
			}
			conf.args = args
			err := command.callback(&conf)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}
