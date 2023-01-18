package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func CreateAnsibleProjectAction(c *cli.Context) error {
	if c.Args().Len() > 0 {
		return errors.New("No arguments is expected, use flags")
	}

	if !c.IsSet("name") {
		return errors.New("A name for the project is required")
	}

	// Create main project directory structure
	if err := createAnsibleProject(c); err != nil {
		return errors.New("Can't create project with provided name. " +
			"Please check if the directory name already exists.")
	}

	// Prompt and create role(s)
	doesUserNeedRoleDirectory(c.String("name"))

	fmt.Println("Project created successfully.")
	return nil
}

func createAnsibleProject(c *cli.Context) error {
	projectName := c.String("name")
	if err := os.Mkdir(fmt.Sprint(projectName), os.ModePerm); err != nil {
		log.Fatal(err)
		return errors.New(fmt.Sprint("Can't create "))
	}

	for _, destination := range mainDirectoryStructure {
		if destination.Type == "directory" {
			if err := os.Mkdir(fmt.Sprint(projectName, destination.Path, "/",
				destination.Name), os.ModePerm); err != nil {
				log.Fatal(err)
				return errors.New("Can't create " + destination.Type + " at " +
					destination.Path + "/" + destination.Name)
			}
		} else {
			pathToFile := fmt.Sprint(projectName, destination.Path, "/")
			fullFileToCreatePath := fmt.Sprint(projectName, destination.Path, "/",
				destination.Name)

			if _, err := os.Stat(pathToFile); os.IsNotExist(err) {
				os.MkdirAll(pathToFile, 0777)
			}

			// Create file
			if _, err := os.OpenFile(fullFileToCreatePath, os.O_RDONLY|os.O_CREATE, 0666); err != nil {
				log.Fatal(err)
				return errors.New("Can't create " + destination.Type + " at " +
					destination.Path + "/" + destination.Name)
			}
		}
	}

	return nil
}

func doesUserNeedRoleDirectory(projectName string) string {
	fmt.Println(projectName + "PROJECT NAME")
	for {
		fmt.Print("Do you want to create Ansible role in this project? Y/n ")

		wantRole, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		wantRoleCaseInsensitive := strings.ToLower(strings.TrimSpace(wantRole))

		if wantRoleCaseInsensitive == "n" {
			return ""
		}

		if wantRoleCaseInsensitive == "y" {
			fmt.Print("Enter the name of role: ")

			roleName, _ := bufio.NewReader(os.Stdin).ReadString('\n')

			roleNameTrimmed := strings.TrimSpace(roleName)

			createRoleDirectory(projectName, roleNameTrimmed)
		}

	}

}

func createRoleDirectory(projectName string, roleName string) error {

	pathToRole := fmt.Sprint(projectName, "/roles/", roleName)

	fmt.Println("Creating role " + roleName + "...")

	if err := os.MkdirAll(pathToRole, os.ModePerm); err != nil {
		log.Fatal(err)
		return errors.New(fmt.Sprint("Can't create directory for role ", roleName))
	}

	for _, destination := range roleDirectoryStructure {
		if destination.Type == "directory" {
			if err := os.Mkdir(fmt.Sprint(projectName, destination.Path, "/", roleName,
				destination.Name), os.ModePerm); err != nil {
				log.Fatal(err)
				return errors.New("Can't create " + destination.Type + " at " +
					destination.Path + "/" + destination.Name)
			}
		} else {
			pathToFile := fmt.Sprint(projectName, destination.Path, "/", roleName)
			fullFileToCreatePath := fmt.Sprint(projectName, destination.Path, "/",
				roleName, destination.Name)

			if _, err := os.Stat(pathToFile); os.IsNotExist(err) {
				os.MkdirAll(pathToFile, 0777)
			}

			// Create file
			if _, err := os.OpenFile(fullFileToCreatePath, os.O_RDONLY|os.O_CREATE, 0666); err != nil {
				log.Fatal(err)
				return errors.New("Can't create " + destination.Type + " at " +
					destination.Path + "/" + destination.Name)
			}
		}
	}

	return nil
}
