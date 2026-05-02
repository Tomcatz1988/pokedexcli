package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	pokecache "https://github.com/Tomcatz1988/pokedexcli/tree/main/internal/pokecache"
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
		if len(words) == 0 {
			continue
		}
		command, exists := reg[words[0]]
		if exists {
			err := command.callback(&conf)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
