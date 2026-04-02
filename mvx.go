package main

import (
	"errors"
	"fmt"
	"os"
)

func parseArgs(args []string) (string, string, error) {
	if len(args) < 2 {
		return "", "", errors.New("usage: mvx <source> <destination>")
	}

	src := args[0]
	dest := args[1]

	return src, dest, nil
}

func handlerError(err error) {
	fmt.Fprintln(os.Stderr, err)
	fmt.Fprintln(os.Stdout, "Try 'mvx --help' for more info.")
	os.Exit(1)
}
