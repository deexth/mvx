package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	HomeDir string
}

type paths struct {
	src  []Source
	dest Dest
}

func (cfg *Config) parseArgs(args []string) (paths, error) {
	if len(args) < 2 {
		return paths{}, errors.New("not enough arguments provided")
	}

	var src []Source

	if len(args) > 2 {
		for _, arg := range args[0 : len(args)-1] {
			source, err := cfg.handlerSource(arg)
			if err != nil {
				return paths{}, err
			}

			src = append(src, source)
		}
	} else {
		source, err := cfg.handlerSource(args[0])
		if err != nil {
			return paths{}, err
		}
		src = append(src, source)
	}

	dest, err := cfg.handlerDestination(args[len(args)-1])
	if err != nil {
		return paths{}, err
	}

	return paths{
		src:  src,
		dest: dest,
	}, nil
}

func cleanPath(path string) string {
	return filepath.Clean(path)
}

func getAbsPath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("absolute path for %s, not found %v", path, err)
	}

	return absPath, nil
}

func handlerError(err error) {
	fmt.Fprintln(os.Stderr, err)
	fmt.Fprintln(os.Stdout, "Try 'mvx --help' for more info.")
	os.Exit(1)
}
