package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	c := &Content{
		text: "hello world",
	}
	command := &AddPrefixCommand{
		prefix: "###",
	}

	command.do(c)
	assert.Equal(t, c.text, "###hello world")

	command.undo(c)
	assert.Equal(t, c.text, "hello world")
}
