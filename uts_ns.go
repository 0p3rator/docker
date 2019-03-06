// UTS Namespace mainly is used to seperate nodename and domainanme

// Enter the following commands to test it both in "sh" shell environment and in a new system terminal shell environment

// "echo $$" output current pid

// "readlink /proc/{pid}/ns/uts" output current UTS namespace

// "hostname -b bird" change the hostname to "bird" in current shell environment
package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
