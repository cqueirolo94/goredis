package command

import (
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
func (cm *CommandMap) Run(input []byte) (string, error) {
	cmdName, args := ParseInput(input)
	cmdName = strings.ToLower(cmdName)

	if cmdName == "exit" {
		return "", errors.New("exit")
	}

	cmd := cm.Commands[cmdName]
	if cmd == nil {
		return "", errors.New("command does not exists")
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
