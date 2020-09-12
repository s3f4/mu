package mu

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommands(t *testing.T) {
	output, err := RunCommands("echo test")
	assert.Nil(t, err)
	assert.Equal(t, []byte("test\n"), output)

	// There is no x folder
	// we are expecting that will throw error
	newOutput, err := RunCommands("cd x")
	// fmt.Printf("%T\n", err)
	// fmt.Printf("%s\n", err.Error())
	assert.Nil(t, newOutput)
	assert.IsType(t, err, &exec.ExitError{})
}
