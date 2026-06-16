package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Pempho-Mackson-Kapulula/pokedex/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	timeout := 5 * time.Second
	cacheIntervalTime := 5 * time.Minute
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(timeout,cacheIntervalTime),
		pokedex: make(map[string]pokeapi.PokemonData),
	}


	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]
		args := cleanedInput[1:]
		if cmd, exists := getCommands()[command]; exists {
			err := cmd.callback(cfg, args...)

			if err != nil {
				fmt.Printf("failed to execute command: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}
