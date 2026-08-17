package cmd

import (
	"fmt"
	"os"
	"strings"
)

func HandleEcho(commandName string, args ...string) {
	fmt.Fprintln(os.Stdout, strings.Join(args, " "))
}
