package command

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that an input with no spaces returns a nil slice as arguments.
func TestEmptyInput(t *testing.T) {
	input := []byte("test_input_no_spaces")
	cmdName, args := ParseInput(input)
	assert.True(t, cmdName == string(input), "the command name wasn't the full input")
	assert.True(t, args == nil, "arguments should be nil")
}

// Test that a correct input of three arguments get parsed correctly
func TestThreeArgsInput(t *testing.T) {
	cmdName := "cmdName"
	arg1 := "arg1"
	arg2 := "arg2"
	arg3 := "arg3"
	input := []byte(fmt.Sprintf("%s %s %s %s", cmdName, arg1, arg2, arg3))

	parsedCmdName, args := ParseInput(input)
	assert.True(t, parsedCmdName == cmdName, fmt.Sprintf("command name wasn't as expected: %s", parsedCmdName))

	assert.True(t, len(args) == 3, fmt.Sprintf("unexpected args len: %d", len(args)))
	assert.True(t, args[0] == arg1, fmt.Sprintf("expected: %s, got: %s", arg1, args[0]))
	assert.True(t, args[1] == arg2, fmt.Sprintf("expected: %s, got: %s", arg2, args[1]))
	assert.True(t, args[2] == arg3, fmt.Sprintf("expected: %s, got: %s", arg3, args[2]))
}
