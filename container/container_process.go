package container

import (
	"os"
	"os/exec"
)

func NewContainerProcess(tty bool, command string) *exec.Cmd {
	cmd := exec.Command(command)
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}
