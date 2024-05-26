package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define an array of valid commands
	validCommands := []string{"exit"}

	// Create a new reader
	reader := bufio.NewReader(os.Stdin)

	for {
		// Print the prompt
		fmt.Fprint(os.Stdout, "$ ")

		// Read the input until newline character
		input, err := reader.ReadString('\n')

		// Check for errors
		if err != nil {
			fmt.Fprintln(os.Stderr, "An error occurred:", err)
			return
		}

		// Trim the newline character from the input
		input = strings.TrimSpace(input)

		words := strings.Fields(input)
		if len(words) == 0 {
			continue // if no words, prompt again
		}

		// Get the first word
		firstWord := words[0]
		// Check if the input is a valid command
		isValidCommand := false
		for _, cmd := range validCommands {
			if firstWord == cmd {
				isValidCommand = true
				break
			}
		}

		// If the command is not valid, print an error message
		if !isValidCommand {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
		} else if isValidCommand && firstWord == "exit" {
			break
		} else {
			fmt.Fprintln(os.Stdout, "Valid command entered:", input)
		}
	}
}
