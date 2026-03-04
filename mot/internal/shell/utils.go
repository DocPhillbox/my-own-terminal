package shell

import (
	"os"
	"os/user"
	"strings"
)

const (
	colorReset = "\033[0m"

	colorGreen  = "\033[32m"
	colorRed    = "\033[31m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorYellow = "\033[33m"
)

func currentUser() string {
	u, err := user.Current()
	if err != nil {
		return "unknown"
	}
	return u.Username
}

func hostname() string {
	h, err := os.Hostname()
	if err != nil {
		return "host"
	}
	return h
}

func currentDir() string {
	wd, err := os.Getwd()
	if err != nil {
		return "?"
	}

	home, err := os.UserHomeDir()
	if err == nil && strings.HasPrefix(wd, home) {
		return "~" + strings.TrimPrefix(wd, home)
	}

	return wd
}
