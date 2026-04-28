package main

import(
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		fmt.Printf("Your command was: %v\n", words[0])
		if words[0] == "exit" {
			break
		}
	}
}
