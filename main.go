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
	paths, err := cfg.parseArgs(args[1:])
	if err != nil {
		handlerError(err)
	}

	fmt.Fprintf(os.Stdout, "source provided is %s and the destination is %s\n", paths.src.loc, paths.dest)

	err = cfg.renameFile(paths.src.loc, paths.dest.loc)
	if err != nil {
		handlerError(err)
	}
}
