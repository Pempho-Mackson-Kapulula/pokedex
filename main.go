package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
			err := cmd.callback()

			if err != nil {
				fmt.Printf("failed to execute command: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command")
		}
	}
}
