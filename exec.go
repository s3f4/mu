package mu

import (
	"os"
	"os/exec"
)

// RunCommands runs multiple commands
func RunCommands(command string) error {
	executable := exec.Command("/bin/sh", "-c", command)
	executable.Stdout = os.Stdout
	executable.Stderr = os.Stderr

	if err := executable.Start(); err != nil {
		return err
	}

	executable.Wait()
	return nil
}
