package main

import (
	"log"
	"os"

	"github.com/mitsuru793/go-packun/commands"
	"github.com/urfave/cli"
)

var (
	cmds  []cli.Command
	types []string
)

func init() {
	cmds = []cli.Command{
		commands.NewAdd(),
		commands.NewInstall(),
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "packun"
	app.Version = "0.0.0"
	app.Usage = "This is cross package manager."

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Commands = cmds

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
