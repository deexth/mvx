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
	paths, err := cli.ParseArgs(args[1:], cfg)
	if err != nil {
		cli.HandlerError(err)
	}

	cfg.PathArgs = paths

	if err = actions(cfg); err != nil {
		cli.HandlerError(err)
	}
}
