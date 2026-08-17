package cmd

import (
	"fmt"
	"os"
	"strings"
)

func HandleCd(command string, args ...string) {
	path := args[0]
	tild := "~"
	if strings.HasPrefix(path, tild) {
		homeDirectory := os.Getenv("HOME")
		path = strings.Replace(path, tild, homeDirectory, 1)
	}

	err := os.Chdir(path)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\r\n", path)
	}
}
