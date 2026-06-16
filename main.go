package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Pempho-Mackson-Kapulula/pokedex/internal/pokeapi"
	"github.com/chzyer/readline"
)

func main() {
	timeout := 5 * time.Second
	cacheIntervalTime := 5 * time.Minute
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(timeout, cacheIntervalTime),
		pokedex:       make(map[string]pokeapi.PokemonData),
	}


	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "." 
	}
	historyFile := filepath.Join(homeDir, ".pokedex_history")

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "Pokedex > ",
		HistoryFile:     historyFile,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		log.Fatalf("failed to initialize readline: %v", err)
	}
	defer rl.Close()

	for {
		
		userInput, err := rl.Readline()
		if err != nil {
			if err == readline.ErrInterrupt { 
				if len(userInput) == 0 {
					break
				}
				continue
			} else if err == io.EOF { 
				break
			}
			fmt.Printf("Error reading input: %v\n", err)
			break
		}

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
