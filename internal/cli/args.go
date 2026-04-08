package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	HomeDir string
}

func ParseArgs(args []string, cfg *Config) (Paths, error) {
	if len(args) < 2 {
		return Paths{}, errors.New("not enough arguments provided")
	}

	var src []Source

	if len(args) > 2 {
		for _, arg := range args[0 : len(args)-1] {
			source, err := handlerSource(arg)
			if err != nil {
				return Paths{}, err
			}

			src = append(src, source)
		}
	} else {
		source, err := handlerSource(args[0])
		if err != nil {
			return Paths{}, err
		}
		src = append(src, source)
	}

	dest, err := handlerDestination(args[len(args)-1])
	if err != nil {
		return Paths{}, err
	}

	return Paths{
		Src:  src,
		Dest: dest,
	}, nil
}

func CleanPath(path string) string {
	return filepath.Clean(path)
}

func GetAbsPath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("absolute path for %s, not found %v", path, err)
	}

	return absPath, nil
}

func HandlerError(err error) {
	fmt.Fprintln(os.Stderr, err)
	fmt.Fprintln(os.Stdout, "Try 'mvx --help' for more info.")
	os.Exit(1)
}
