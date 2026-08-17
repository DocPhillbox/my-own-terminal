package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func HandleCommand(command string) {
	commandName, args := ParseCommand(command)
	hasOperand, operator := HasOperand(args)
	var operands []string
	if hasOperand {
		args, operands = HandleRedirection(args)
	}
	found, _ := IsExecutable(commandName)
	if handler, ok := commands[commandName]; ok {
		if hasOperand {
			bultinWithOperand(handler, commandName, args, operands, operator)
		} else {
			handler(commandName, args...)
		}
	} else if found {
		cmd := exec.Command(commandName, args...)
		if hasOperand {
			externalWithOperand(cmd, operands, operator)
		} else {
			external(cmd, commandName)
		}
	} else {
		fmt.Printf("%s: command not found\r\n", commandName)
	}
}

func bultinWithOperand(handler func(string, ...string), commandName string, args []string, operands []string, operator int) {
	var file *os.File
	var err error
	switch operator {
	case 1, 2:
		file, err = os.Create(operands[0])
		if err != nil {
			fmt.Printf("Error opening file for redirection: %v\r\n", err)
			return
		}
	case 3, 4:
		file, err = os.OpenFile(operands[0], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening file for redirection: %v\r\n", err)
			return
		}
	}
	defer file.Close()
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	switch operator {
	case 1:
		os.Stdout = file
	case 2:
		os.Stderr = file
	case 3:
		os.Stdout = file
	case 4:
		os.Stderr = file
	}
	handler(commandName, args...)
	switch operator {
	case 1:
		os.Stdout = oldStdout
	case 2:
		os.Stderr = oldStderr
	case 3:
		os.Stdout = oldStdout
	case 4:
		os.Stderr = oldStderr
	}
}

func external(cmd *exec.Cmd, commandName string) {
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing %s: %v\r\n", commandName, err)
	}
	fmt.Print(string(cmdOutput))
}

func externalWithOperand(cmd *exec.Cmd, operands []string, operator int) {
	var file *os.File
	var err error
	switch operator {
	case 1, 2:
		file, err = os.Create(operands[0])
		if err != nil {
			fmt.Printf("Error opening file for redirection: %v\r\n", err)
			return
		}
	case 3, 4:
		file, err = os.OpenFile(operands[0], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening file for redirection: %v\r\n", err)
			return
		}
	}
	defer file.Close()
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	switch operator {
	case 1:
		cmd.Stdout = file
		cmd.Stderr = os.Stderr
	case 2:
		cmd.Stdout = os.Stdout
		cmd.Stderr = file
	case 3:
		cmd.Stdout = file
		cmd.Stderr = os.Stderr
	case 4:
		cmd.Stdout = os.Stdout
		cmd.Stderr = file
	}
	cmd.Run()
	switch operator {
	case 1:
		os.Stdout = oldStdout
		os.Stderr = oldStderr
	case 2:
		os.Stdout = oldStdout
		os.Stderr = oldStderr
	case 3:
		os.Stdout = oldStderr
		os.Stderr = oldStderr
	case 4:
		os.Stdout = oldStdout
		os.Stderr = oldStderr
	}
}
