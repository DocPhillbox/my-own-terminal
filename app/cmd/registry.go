package cmd

import (
	"os"
)

var commands = map[string]func(string, ...string){
	"exit": func(_ string, _ ...string) { os.Exit(0) },
	"echo": HandleEcho,
	"pwd":  HandlePwd,
	"cd":   HandleCd,
}

func init() {
	commands["type"] = HandleType
}
