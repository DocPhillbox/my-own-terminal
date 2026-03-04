package main

import (
	"fmt"
	"os"

	"github.com/DocPhillbox/mot/internal/shell"
)

func main() {
	if err := shell.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
