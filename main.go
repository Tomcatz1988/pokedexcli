package main

import (
	"bufio"
	"fmt"
	"os"
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
