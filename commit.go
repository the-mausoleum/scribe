package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseArgs(args []string) (Commit, error) {
	if len(args) < 2 {
		return Commit{}, errors.New("Not enough arguments.")
	}

	var project string = args[0]
	hours, err := strconv.Atoi(args[1])

	if err != nil {
		return Commit{}, err
	}

	var message string

	if len(args) > 2 {
		message = strings.Join(args[2:], " ")
	} else {
		message = ""
	}

	return Commit{project, hours, message}, nil
}

func DoCommit(args []string) {
	commit, err := parseArgs(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(commit)

	CreateCommit(commit)
}
