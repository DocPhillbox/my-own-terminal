package shell

import (
	"fmt"
	"os"
	"strings"
)

type Shell struct {
	lastStatus int
	history    []string
}

func Run() error {
	s := Shell{}
	handleSignals(&s)

	for {
		fmt.Print(s.prompt())

		line, err := s.readLine()
		if err != nil {
			return err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		err = parseCommand(&s, line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			s.lastStatus = 1
		} else {
			s.lastStatus = 0
		}

		const maxHistory = 1000

		if len(s.history) >= maxHistory {
			s.history = s.history[1:]
		}
		if len(s.history) == 0 || s.history[len(s.history)-1] != line {
			s.history = append(s.history, line)
		}
	}
}
