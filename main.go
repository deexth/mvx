package main

import (
	"fmt"
	"os"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to get the user's home directory: %v", err)
		os.Exit(1)
	}

	cfg := &Config{
		HomeDir: homeDir,
	}

	args := os.Args
	src, dest, err := parseArgs(args[1:], cfg)
	if err != nil {
		handlerError(err)
	}

	fmt.Fprintf(os.Stdout, "source provided is %s and the destination is %s\n", src, dest)
}
