package main

import (
	"docker/mydocker/container"
	log "github.com/Sirupsen/logrus"
	"os"
)

/*
	Start() here actually start to fork "cmd" created by "NewParentProcess()"
		1. Clone a new process of which namespace is seperated with external environment
		2. In the new process, fork "/proc/self/exec", send "init" args, and fork "initCommand" we have writed
*/
func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
