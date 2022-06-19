package command

import "errors"

// PING command
func Ping() *Command {
	return &Command{
		Name:         "PING",
		ExpectedArgs: 0,
		Fn: func(args ...string) (string, error) {
			if len(args) != 0 {
				return "", errors.New("no arguments required for this command")
			}
			return "PONG", nil
		},
	}
}
