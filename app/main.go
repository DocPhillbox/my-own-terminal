package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/sys/unix"
	"golang.org/x/term"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
)

func main() {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error setting raw mode:", err)
		os.Exit(1)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	termios, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		log.Fatal(err)
	}

	termios, errTermios := unix.IoctlGetTermios(fd, unix.TCGETS)
	if errTermios != nil {
		log.Fatal(err)
	}

	termios.Oflag |= unix.OPOST | unix.ONLCR

	if err := unix.IoctlSetTermios(fd, unix.TCSETS, termios); err != nil {
		log.Fatal(err)
	}

	prompt := "$ "
	fmt.Print(prompt)
	var input []rune
	var tab byte = '\t'
	var backspace byte = 127
	var escape byte = 27
	var bell byte = '\a'
	tabbed := false
	buf := make([]byte, 1)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		ch := buf[0]
		if ch != tab {
			tabbed = false
		}
		if ch == '\n' || ch == '\r' {
			fmt.Print("\r\n")
			command := strings.TrimSpace(string(input))
			cmd.HandleCommand(command)
			input = input[:0]
			fmt.Print("\r$ ")
			continue
		}
		if ch == tab {
			partialCommand := string(input)
			commands := cmd.Autocomplete(string(input))
			commands = cmd.RemoveDuplicates(commands)
			commandsSize := len(commands)
			if commandsSize == 1 {
				command := commands[0]
				for range len(partialCommand) {
					input = cmd.RemoveChar(input)
				}
				fmt.Print(command + " ")
				for _, char := range command {
					input = append(input, char)
				}
				input = append(input, ' ')
			} else {
				if tabbed {
					fmt.Println()
					cmd.ShowCommands(commands)
					cmd.ShowPrompt(prompt, input)
				} else {
					os.Stdout.Write([]byte{bell})
				}
			}
			tabbed = true
			continue
		}
		if ch == 3 {
			fmt.Println()
			os.Exit(0)
		}
		if ch == backspace {
			if len(input) > 0 {
				input = cmd.RemoveChar(input)
			}
			continue
		}
		if ch == escape {
			os.Stdin.Read(buf)
			os.Stdin.Read(buf)
			continue
		}
		input = append(input, rune(ch))
		fmt.Printf("%c", ch)
	}
}
