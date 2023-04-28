package cmd_test

import (
	"testing"

	"github.com/brpaz/copier-run/cmd"
	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	command := cmd.NewVersionCmd()

	assert.Equal(t, "version", command.Use)
	assert.Nil(t, command.Execute())
	assert.NotNil(t, command)
}
