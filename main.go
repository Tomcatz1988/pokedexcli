package main

import (
	"bufio"
	"fmt"
	"os"

	pokecache "internal/pokecache"
)

type config struct {
	Next string
	Previous string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	reg := getCommandsRegistry()
	conf := config{
		locationURL,
		locationURL,
	}
	cache := pokecache.NewCache(cacheDuration)

	for {
		fmt.Print("Pokedex > ")
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
			err := command.callback(&conf, &cache, args)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command")
		}
		fmt.Println("")
	}
}
