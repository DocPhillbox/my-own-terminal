package shell

import (
	"os"
	"os/exec"
	"strings"
)

func parseCommand(s *Shell, input string) error {
	commands := strings.SplitSeq(input, ";")

	for command := range commands {
		cmd := strings.TrimSpace(command)
		if cmd == "" {
			continue
		}

		err := parseOperators(cmd, s)
		if err != nil {
			s.lastStatus = 1
		} else {
			s.lastStatus = 0
		}
	}
	return nil
}

func parseOperators(cmd string, s *Shell) error {
	if strings.Contains(cmd, "&&") {
		parts := strings.SplitN(cmd, "&&", 2)

		err := executeSingle(strings.TrimSpace(parts[0]), s)
		if err == nil {
			return executeSingle(strings.TrimSpace(parts[1]), s)
		}
		return err
	}

	if strings.Contains(cmd, "||") {
		parts := strings.SplitN(cmd, "||", 2)

		err := executeSingle(strings.TrimSpace(parts[0]), s)
		if err != nil {
			return executeSingle(strings.TrimSpace(parts[1]), s)
		}
		return nil
	}

	return executeSingle(cmd, s)
}

func executeSingle(cmd string, s *Shell) error {
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return nil
	}

	name := parts[0]
	args := parts[1:]

	switch name {
	case "exit":
		os.Exit(0)

	case "history":
		s.builtinHistory()
		return nil

	case "cd":
		return builtinCd(args)

	default:
		return runExternal(name, args)
	}
	return nil
}

func runExternal(name string, args []string) error {
	command := exec.Command(name, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	return command.Run()
}
