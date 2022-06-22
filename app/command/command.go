package command

import (
	"bytes"
	"errors"
	"strings"
)

// A map that holds all available functions tied to their name
type CommandMap struct {
	Commands map[string]*Command
}

// A command has a name and a function associated.
type Command struct {
	Name         string
	ExpectedArgs int
	Fn           func(...string) (string, error)
}

// Returns the result of executing the function of the command passed with its arguments
func (cm *CommandMap) Run(byteInput []byte) (string, error) {
	trimmedBytes := bytes.Trim(byteInput, "\x00\n")

	cmdName, args := ParseInput(string(trimmedBytes))

	if strings.ToLower(cmdName) == "exit" {
		return "", errors.New("exit")
	}

	cmd, err := cm.Commands[cmdName]
	if !err {
		return "Command doesn't exist", nil
	}
	return cmd.Fn(args...)
}

// Creates and returns a new CommandMap pointer, with the commands saved in it.
func New() *CommandMap {
	commandMap := &CommandMap{
		Commands: make(map[string]*Command),
	}

	// Ping command
	pingCmd := Ping()
	commandMap.Commands[pingCmd.Name] = pingCmd

	return commandMap
}
