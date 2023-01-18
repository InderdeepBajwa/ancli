package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "The CLI for Ansible that you always needed."
	app.Usage = "Create Ansible projects, add new roles, and manage dependencies."

	app.Commands = []*cli.Command{
		{
			Name:        "create",
			HelpName:    "create",
			Action:      CreateAnsibleProjectAction,
			ArgsUsage:   ` `,
			Usage:       "Creates a new Ansible project in the current directory",
			Description: "Create a new Ansible project",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "Name of the project",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal((err))
	}
}
