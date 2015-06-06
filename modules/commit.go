package modules

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type commit struct {
	project string
	hours   int
	message string
}

func parseArgs(args []string) (commit, error) {
	if len(args) < 2 {
		return commit{}, errors.New("Not enough arguments.")
	}

	var project string = args[0]
	hours, err := strconv.Atoi(args[1])

	if err != nil {
		return commit{}, err
	}

	var message string

	if len(args) > 2 {
		message = strings.Join(args[2:], " ")
	} else {
		message = ""
	}

	return commit{project, hours, message}, nil
}

func Commit(args []string) {
	commit, err := parseArgs(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(commit)
}
