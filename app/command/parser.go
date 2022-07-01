package command

import (
	"bytes"
	"strings"
)

// Given an input, it returns the command name in uppercase, and the arguments splitted
func ParseInput(input []byte) (string, []string) {
	trimmedInput := bytes.Trim(input, "\x00\n")

	words := strings.Split(string(trimmedInput), " ")
	if len(words) == 1 {
		return words[0], nil
	}

	return words[0], words[1:]
}
