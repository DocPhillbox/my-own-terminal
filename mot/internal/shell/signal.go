package shell

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handleSignals(s *Shell) {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		for range sigChan {
			fmt.Println()
			fmt.Print(s.prompt())
		}
	}()
}
