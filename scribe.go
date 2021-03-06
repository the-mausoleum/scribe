package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	Version = "1.0.0"
	Usage   = `
Usage:

        %s command [arguments]
`
)

var (
	versionFlag = flag.Bool("version", false, "Show current version information.")
	helpFlag    = flag.Bool("help", false, "Show the help.")
)

func init() {
	flag.BoolVar(versionFlag, "v", false, "Get the current version.")
	flag.BoolVar(helpFlag, "h", false, "Show the help.")
}

func main() {
	flag.Parse()

	switch {
	case *versionFlag:
		fmt.Printf("%s %s\n", os.Args[0], Version)
		os.Exit(0)
	case *helpFlag:
		fmt.Printf(Usage, os.Args[0])
		os.Exit(0)
	}

	Connect()

    args := os.Args[2:]

	switch strings.ToLower(os.Args[1]) {
	case "commit":
		DoCommit(args)
    case "rename":
        DoRename(args)
	case "delete":
		DoDelete(args)
	default:
		fmt.Printf(Usage, os.Args[0])
		os.Exit(0)
	}
}
