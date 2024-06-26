package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const usage = "minidocker is a simple container runtime implementation."

func main() {
	app := cli.NewApp()
	app.Name = "minidocker"
	app.Usage = usage

	app.Commands = []cli.Command{
		runCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
