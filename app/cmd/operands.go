package cmd

import (
	"slices"
	"strings"
)

var operands = []string{">", "1>", "2>", ">>", "1>>", "2>>"}

func HasOperand(args []string) (bool, int) {
	hasOperand := false
	operator := 0
	for _, part := range args {
		part = strings.TrimSpace(part)
		if slices.Contains(operands, part) {
			hasOperand = true
			switch part {
			case operands[0], operands[1]:
				operator = 1
			case operands[2]:
				operator = 2
			case operands[3], operands[4]:
				operator = 3
			case operands[5]:
				operator = 4
			}
		}
	}
	return hasOperand, operator
}

func HandleRedirection(args []string) ([]string, []string) {
	var cmdPart []string
	var filePart []string
	for i, part := range args {
		if slices.Contains(operands, part) {
			cmdPart = args[0:i]
			filePart = args[i+1:]
			return cmdPart, filePart
		}
	}
	return cmdPart, filePart
}
