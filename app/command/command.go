package command

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
func (cm *CommandMap) Run(input string) (string, error) {
	// TODO: validate the string passed.

	cmdName, args := ParseInput(input)

	return cm.Commands[cmdName].Fn(args...)
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
