package container

import (
	"os"
	"os/exec"
	"syscall"
)

func NewContainerProcess(tty bool, command string) *exec.Cmd {
	cmd := exec.Command(command)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}
