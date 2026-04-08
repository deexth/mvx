package main

import (
	"fmt"
	"os"

	"github.com/deexth/mvx/internal/cli"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to get the user's home directory: %v", err)
		os.Exit(1)
	}

	cfg := &cli.Config{
		HomeDir: homeDir,
	}

	args := os.Args
	Paths, err := cfg.parseArgs(args[1:])
	if err != nil {
		handlerError(err)
	}

	fmt.Fprintf(os.Stdout, "source provided is %s and the destination is %s\n", Paths.src, paths.dest)

	err = cfg.renameFile(Paths.src, Paths.dest.loc)
	if err != nil {
		handlerError(err)
	}
}
