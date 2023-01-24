package main

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
)

func CreateAnsibleRoleAction(c *cli.Context) error {
	if c.Args().Len() > 0 {
		return errors.New("No arguments is expected, use flags")
	}

	if !c.IsSet("name") {
		return errors.New("A name for the project is required")
	}

	// Create Ansible role in specified directory
	createRoleDirectory(c.String("project"), c.String("name"))

	fmt.Println("Role created successfully.")
	return nil
}
