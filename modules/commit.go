package modules

import (
	"fmt"
    "os"
    "strconv"
)

type commit struct {
    project string
    hours int
    message string
}

func parseArgs(args []string) commit {
    hours, err := strconv.Atoi(args[1])
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }

    return commit{args[0], hours, args[2]}
}

func Commit(args []string) {
    commit := parseArgs(args)

    fmt.Println(commit)
}
