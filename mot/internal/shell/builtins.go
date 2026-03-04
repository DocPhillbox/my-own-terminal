package shell

import (
	"fmt"
	"os"
)

func builtinCd(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("cd: missing argument")
	}

	return os.Chdir(args[0])
}

func (s *Shell) builtinHistory() {
	for i, cmd := range s.history {
		fmt.Printf("%5d  %s\n", i+1, cmd)
	}
}
