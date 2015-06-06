package main

import (
	"errors"
	"fmt"
	"os"
)

func DoDelete(args []string) {
	if len(args) < 1 {
		fmt.Println(errors.New("Must specify a project name."))
		os.Exit(0)
	}

	var project string = args[0]

	if project == "" {
		fmt.Println(errors.New("Must specify a project name."))
		os.Exit(0)
	}

	DeleteProject(project)
}

func DoRename(args []string) {
	if len(args) < 2 {
		fmt.Println(errors.New("Missing required arguments."))
		os.Exit(0)
	}

	var project, newName string = args[0], args[1]
	if project == "" {
		fmt.Println(errors.New("Must specify a project name."))
		os.Exit(0)
	}

	if newName == "" {
		fmt.Println(errors.New("Must specify a new project name."))
		os.Exit(0)
	}

	RenameProject(project, newName)
}
