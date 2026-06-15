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

	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
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
		if cmd, exists := getCommands()[command]; exists {
			err := cmd.callback(cfg)

			if err != nil {
				fmt.Printf("failed to execute command: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command")
		}
	}
}
