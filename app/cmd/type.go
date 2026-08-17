package cmd

import (
	"fmt"
)

func HandleType(command string, args ...string) {
	name := args[0]
	if _, ok := commands[name]; ok {
		fmt.Printf("%s is a shell builtin\r\n", name)
	} else {
		found, cmdPath := IsExecutable(name)
		if found {
			fmt.Printf("%s is %s\r\n", name, cmdPath)
		} else {
			fmt.Printf("%s not found\r\n", name)
		}
	}
}
