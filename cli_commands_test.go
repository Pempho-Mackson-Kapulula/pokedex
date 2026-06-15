package main

import "testing"

func TestGetCommands(t *testing.T) {
	// 1. Get the map from your function
	commands := getCommands()

	// 2. Define the exact keys we expect to exist
	expectedKeys := []string{"exit", "help", "map", "mapb"}

	// 3. Verify the map length matches expected keys
	if len(commands) != len(expectedKeys) {
		t.Errorf("Expected %d commands, but got %d", len(expectedKeys), len(commands))
	}

	// 4. Validate each expected command structure
	for _, key := range expectedKeys {
		cmd, exists := commands[key]
		if !exists {
			t.Errorf("Expected command key '%s' was not found in the map", key)
			continue
		}

		// Verify the internal name matches the map key
		if cmd.name != key {
			t.Errorf("Expected internal name '%s', got '%s'", key, cmd.name)
		}

		// Verify the description field is not empty
		if cmd.description == "" {
			t.Errorf("Description for command '%s' should not be empty", key)
		}

		// Verify the callback function pointer is properly assigned
		if cmd.callback == nil {
			t.Errorf("Callback function for command '%s' is nil", key)
		}
	}
}

func TestCommandHelpExecution(t *testing.T) {
	// We can safely invoke commandHelp because it only prints to console
	err := commandHelp(&config{})
	if err != nil {
		t.Errorf("commandHelp returned an unexpected error: %v", err)
	}
}
