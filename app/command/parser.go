package command

import (
	"strings"
)

// Given an input, it returns the command name in uppercase, and the arguments splitted
func ParseInput(input string) (string, []string) {
	inputTrimmed := strings.TrimSuffix(input, "\n")

	words := strings.Split(inputTrimmed, " ")
	if len(words) == 1 {
		return words[0], nil
	}

	return words[0], words[1:]
}
