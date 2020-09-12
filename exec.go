package mu

import (
	"os/exec"
)

// RunCommands runs multiple commands
func RunCommands(command string) ([]byte, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		return nil, err
	}

	return output, nil
}
