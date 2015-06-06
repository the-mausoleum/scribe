package main

import (
	"flag"
	"fmt"
	"os"
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
		fallthrough
	default:
		fmt.Printf(Usage, os.Args[0])
		os.Exit(0)
	}
}
