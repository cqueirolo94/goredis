package command

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that checks that the ping command returns PONG and nil error when no args where passed
func TestPingWithNoArguments(t *testing.T) {
	pingCmd := Ping()
	res, err := pingCmd.Fn()
	assert.True(t, res == "PONG", "PONG wasn't returned")
	assert.True(t, err == nil, fmt.Sprintf("err wasn't nil: %s", err))
}

func TestPingWithArguments(t *testing.T) {
	pingCmd := Ping()
	res, err := pingCmd.Fn("Arg 1", "Arg 2")
	assert.True(t, res == "", "empty string expected")
	assert.True(t, err != nil, "err was nil, expected error message")
}
