package main

import (
	"testing"
)


func TestCleanInput(t *testing.T){
	cases := []struct {
	input    string
	expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: " This is a Test string",
			expected: []string{"this", "is", "a", "test", "string"},
		},
	}

	for _, c := range cases {
	actual := cleanInput(c.input)
	// Check the length of the actual slice against the expected slice

	limit := len(actual)
	if len(c.expected) < limit {
    limit = len(c.expected)
	}

	if len(actual) != len(c.expected){
				// print an error message if length doesn't match
				t.Errorf("the length of the words do not match")
				
	}


	for i := 0; i < limit; i++{
		
		word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if expectedWord != word{
				t.Errorf("expected: %v, got: %v", expectedWord,word)
			}
		}	
	}
}