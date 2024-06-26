package main

import (
	"os"

	"minidocker/container"

	log "github.com/sirupsen/logrus"
)

func Run(tty bool, cmd string) {
	parent := container.NewContainerProcess(tty, cmd)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	_ = parent.Wait()
	os.Exit(-1)
}
