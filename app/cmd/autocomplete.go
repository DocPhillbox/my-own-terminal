package cmd

import (
	"fmt"
	"slices"
	"strings"
)

func Autocomplete(partialCommand string) []string {
	var foundCommands []string
	for _, executable := range ListExecutable() {
		if strings.HasPrefix(executable, partialCommand) {
			foundCommands = append(foundCommands, executable)
		}
	}
	for commandName := range commands {
		if strings.HasPrefix(commandName, partialCommand) {
			foundCommands = append(foundCommands, commandName)
		}
	}
	return foundCommands
}

func ShowCommands(commands []string) {
	space := "  "
	slices.Sort(commands)
	for i, command := range commands {
		if i == len(commands)-1 {
			fmt.Print(command)
		} else {
			fmt.Print(command + space)
		}
	}
	fmt.Println()
}

func ShowPrompt(prompt string, input []rune) {
	fmt.Print(prompt)
	for _, ch := range input {
		fmt.Printf("%c", ch)
	}
}

func RemoveChar(input []rune) []rune {
	fmt.Print("\b \b")
	return input[:len(input)-1]
}
