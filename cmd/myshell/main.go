package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	validCommands := []string{"exit", "echo", "type"}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "An error occurred:", err)
			return
		}

		input = strings.TrimSpace(input)

		words := strings.Fields(input)
		if len(words) == 0 {
			continue
		}

		firstWord := words[0]
		args := words[1:]

		isValidCommand := false
		for _, cmd := range validCommands {
			if firstWord == cmd {
				isValidCommand = true
				break
			}
		}

		if !isValidCommand {
			err = execCommand(firstWord, args)
			if err != nil {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", firstWord)
			}
		} else if firstWord == "exit" {
			break
		} else if firstWord == "echo" {
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
		} else if firstWord == "type" {
			typeShell(args, validCommands)
		} else {
			fmt.Fprintln(os.Stdout, "Valid command entered:", input)
		}
	}
}

func typeShell(args []string, validCommands []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stdout, "type: missing operand")
		return
	}

	arg := args[0]
	isBuiltin := false
	for _, cmd := range validCommands {
		if arg == cmd {
			isBuiltin = true
			break
		}
	}

	if isBuiltin {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", arg)
	} else {
		paths := strings.Split(os.Getenv("PATH"), ":")
		if path, ok := findExecutable(arg, paths); ok {
			fmt.Printf("%s is %s\n", arg, path)
		} else {
			fmt.Fprintf(os.Stdout, "%s not found\n", arg)
		}
	}
}

func findExecutable(name string, paths []string) (string, bool) {
	for _, path := range paths {
		fullpath := filepath.Join(path, name)
		if _, err := os.Stat(fullpath); err == nil {
			return fullpath, true
		}
	}
	return "", false
}

func execCommand(name string, args []string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
