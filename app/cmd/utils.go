package cmd

import (
	"os"
	"slices"
	"strings"
)

func ListExecutable() []string {
	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")
	var exectuables []string
	for _, p := range paths {
		entries, err := os.ReadDir(p)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			exectuables = append(exectuables, entry.Name())
		}
	}
	return exectuables
}

func IsExecutable(name string) (bool, string) {
	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")
	found := false
	foundPath := ""
	for _, p := range paths {
		cmdPath := p + "/" + name
		if CheckExecutionPermsission(cmdPath) {
			found = true
			foundPath = cmdPath
			break
		}
	}
	return found, foundPath
}

func CheckExecutionPermsission(path string) bool {
	hasPermission := false
	if info, err := os.Stat(path); err == nil {
		if info.Mode().IsRegular() && info.Mode().Perm()&0111 != 0 {
			hasPermission = true
		}
	}
	return hasPermission
}

func ParseCommand(commandText string) (string, []string) {
	var command strings.Builder
	var args strings.Builder
	isCommand := true
	isQuote := false
	isSingleQuote := false
	isDoubleQuote := false
	for _, char := range commandText {
		if char == '\'' && !isDoubleQuote && isCommand {
			isSingleQuote = !isSingleQuote
			isQuote = !isQuote
		}
		if char == '"' && !isSingleQuote && isCommand {
			isDoubleQuote = !isDoubleQuote
			isQuote = !isQuote
		}
		if char == ' ' && !isQuote && isCommand {
			isCommand = false
			continue
		}
		if isCommand {
			command.WriteRune(char)
		} else {
			args.WriteRune(char)
		}
	}
	return ParseQuotesCommand(command.String()), ParseQuotes(args.String())
}

func ParseQuotesCommand(arg string) string {
	var text strings.Builder
	singleQuote := '\''
	doubleQuote := '"'
	backslash := '\\'
	space := ' '
	specialChar := []rune{'"', '\\', '$', '`', '\n'}
	spaced := false
	quotes := false
	singleQuotes := false
	doubleQuotes := false
	backslashed := false
	for _, char := range arg {
		if backslashed && !quotes {
			backslashed = false
			text.WriteRune(char)
			continue
		}
		if backslashed && doubleQuotes {
			backslashed = false
			if slices.Contains(specialChar, char) {
				text.WriteRune(char)
			} else {
				text.WriteRune(backslash)
				text.WriteRune(char)
			}
			continue
		}
		if char == doubleQuote && !singleQuotes {
			doubleQuotes = !doubleQuotes
			quotes = !quotes
			continue
		}
		if char == singleQuote && !doubleQuotes {
			singleQuotes = !singleQuotes
			quotes = !quotes
			continue
		}
		if !quotes && char != space {
			spaced = false
			if char == backslash && !backslashed {
				backslashed = true
			} else {
				text.WriteRune(char)
			}
		} else if !quotes && char == space && !spaced {
			spaced = true
			text.WriteRune(char)
		}
		if quotes {
			spaced = false
			if char == backslash && !backslashed && doubleQuotes {
				backslashed = true
			} else {
				text.WriteRune(char)
			}
		}
	}
	return text.String()
}

func ParseQuotes(arg string) []string {
	var args []string
	var text strings.Builder
	singleQuote := '\''
	doubleQuote := '"'
	backslash := '\\'
	space := ' '
	specialChar := []rune{'"', '\\', '$', '`', '\n'}
	spaced := false
	quotes := false
	singleQuotes := false
	doubleQuotes := false
	backslashed := false
	for _, char := range arg {
		if backslashed && !quotes {
			backslashed = false
			text.WriteRune(char)
			continue
		}
		if backslashed && doubleQuotes {
			backslashed = false
			if slices.Contains(specialChar, char) {
				text.WriteRune(char)
			} else {
				text.WriteRune(backslash)
				text.WriteRune(char)
			}
			continue
		}
		if char == doubleQuote && !singleQuotes {
			doubleQuotes = !doubleQuotes
			quotes = !quotes
			spaced = false
			continue
		}
		if char == singleQuote && !doubleQuotes {
			singleQuotes = !singleQuotes
			quotes = !quotes
			spaced = false
			continue
		}
		if !quotes && char != space {
			spaced = false
			if char == backslash && !backslashed {
				backslashed = true
			} else {
				text.WriteRune(char)
			}
		} else if !quotes && char == space && !spaced {
			args = append(args, text.String())
			text.Reset()
			spaced = true
		}
		if quotes {
			spaced = false
			if char == backslash && !backslashed && doubleQuotes {
				backslashed = true
			} else {
				text.WriteRune(char)
			}
		}
	}
	args = append(args, text.String())
	return args
}

func RemoveDuplicates(list []string) []string {
	seen := make(map[string]struct{})
	result := []string{}

	for _, v := range list {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
