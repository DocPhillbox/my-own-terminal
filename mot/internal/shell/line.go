package shell

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func (s *Shell) readLine() (string, error) {
	fd := int(os.Stdin.Fd())

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return "", err
	}
	defer term.Restore(fd, oldState)

	var line []rune
	historyIndex := len(s.history)

	buf := make([]byte, 1)

	for {
		os.Stdin.Read(buf)
		b := buf[0]

		// ENTER
		if b == '\r' || b == '\n' {
			fmt.Print("\r\n")
			return string(line), nil
		}

		// BACKSPACE
		if b == 127 {
			if len(line) > 0 {
				line = line[:len(line)-1]
				fmt.Print("\b \b")
			}
			continue
		}

		// ESC
		if b == 27 {
			os.Stdin.Read(buf)
			if buf[0] == '[' {
				os.Stdin.Read(buf)

				switch buf[0] {
				case 'A': // ↑
					if historyIndex > 0 {
						historyIndex--
						line = []rune(s.history[historyIndex])
						redrawLine(line, s)
					}
				case 'B': // ↓
					if historyIndex < len(s.history)-1 {
						historyIndex++
						line = []rune(s.history[historyIndex])
						redrawLine(line, s)
					} else {
						historyIndex = len(s.history)
						line = []rune{}
						redrawLine(line, s)
					}
				}
			}
			continue
		}

		line = append(line, rune(b))
		fmt.Printf("%c", b)
	}
}

func redrawLine(line []rune, s *Shell) {
	fmt.Print("\r\033[K")
	fmt.Print(s.prompt())
	fmt.Print(string(line))
}
