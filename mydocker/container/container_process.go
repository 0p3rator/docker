package container

import (
	"syscall"
	"os/exec"
	"os"
)
/*
Here is in parent process.

1. In fork "/proc/self/exe", "/proc/self" is current parent process environment, "exec" is to fork parent process itself, By this way init new created process.

2. "args..." is paramenters, "init" in them is the first arg which passed to this new created process. It will fork "initCommand" later.

3. Clone args following is to fork new process using namespace to seperate it with external environment.

4. If user used "-ti" arg, we need to redirect input and output stream to standard stream.
*/
func NewParentProcess(tty bool, command string) *exec.Cmd {
	args := []string{"init", command}
	cmd := exec.Command("/proc/self/exe", args...)
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
		syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
    }
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}
