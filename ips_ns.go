package main

/*
	1, " ipcs -q " check current ipc message queue
	2, " ipcmk -Q " create a new message queue
	3, " ipcs -q " check current ipc message queue again you will find a message queue created just now

	4, open a new "sh" shell environment by "go run ips_ns.go"
	5, repeat step 1, 2, 3

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
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
