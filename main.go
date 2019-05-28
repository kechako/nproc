package main

import (
	"fmt"
	"os"
	"runtime"
)

const (
	appName    = "nproc"
	appVersion = "1.0.0"
)

func printVersion() {
	fmt.Printf(`%s %s
Copyright (C)2019 Ryosuke Akiyama
`, appName, appVersion)
}

func printHelp() {
	fmt.Printf(`Usage: %s [OPTION]...
Print the number of logical CPUs usable by the current process.

      --help     display this help and exit
      --version  output version information and exit
`, appName)
}

func printError(opt string) {
	fmt.Fprintf(os.Stderr, `%s: unrecognized option '%s'
Try '%s --help' for more information.
`, appName, opt, appName)
}

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(runtime.NumCPU())
	} else {
		opt := args[0]
		switch opt {
		case "--help":
			printHelp()
		case "--version":
			printVersion()
		default:
			printError(opt)
			os.Exit(1)
		}
	}
}
