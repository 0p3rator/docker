package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"docker/mydocker/container"
)

//Here defines Flags of runCommand, its function is like runtime command "--" to specify args
var runCommand = cli.Command{
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit
			mydocker run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	/*
	Here is the real function which "run" command really excute
		1. check args whether contains command or not
		2. get command which user specified
		3. fork "Run" function to init container
	*/
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		Run(tty, cmd)
		return nil
	},
}

//Here defines specific opt for initCommand, initCommand is internal function, can not be fork externally.
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	/*
	1. get args of command from user
	2. excute container's init progress 
	*/
	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		cmd := context.Args().Get(0)
		log.Infof("command %s", cmd)
		//err := container.RunContainerInitProcess(cmd, nil)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}
