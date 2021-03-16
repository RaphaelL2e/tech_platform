package main

import (
	"github.com/urfave/cli/v2"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "tech_platform"
	app.Flags = flags
	app.Action = server

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
