package container

import (
	"os"
	"syscall"
	"github.com/Sirupsen/logrus"
)

/*
	"init()" function here is excuted in container process, in other words, container process has been created since code run to here. This is the first process of container.
	Use "mount" to mount "proc" file system, and "ps" command will use it to check resource use of current process.
*/
func RunContainerInitProcess(command string, args []string) error {
	logrus.Infof("command %s", command)

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}
