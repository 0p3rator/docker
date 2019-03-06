package main

/*
	1. in "sh" shell
			1, "ls /proc" to show current /proc file content
			2, " mount -t proc proc /proc" to mount proc to /proc with type proc
			3, "ls /proc" to show current /proc file content, you will find result number decrease a lot
	2, in system shell "ls /proc" check the difference with "sh" shell result


*/

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
