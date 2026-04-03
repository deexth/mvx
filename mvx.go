package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	HomeDir string
}

func parseArgs(args []string, cfg *Config) (string, string, error) {
	if len(args) < 2 {
		return "", "", errors.New("usage: mvx <source> <destination>")
	}

	for i := range args {
		if !strings.HasPrefix(args[i], "~") {
			continue
		}

		args[i] = strings.Split(args[i], "~")[1]
	}

	src := filepath.Clean(args[0])
	dest := filepath.Clean(args[1])

	err := cfg.parseSrc(src)
	if err != nil {
		return "", "", err
	}

	err = cfg.parseDest(dest)
	if err != nil {
		return "", "", err
	}

	return src, dest, nil
}

func (cfg *Config) parseSrc(path string) error {
	newPath := filepath.Join(cfg.HomeDir, path)

	ok := filepath.IsAbs(newPath)
	if !ok {
		return fmt.Errorf("no such file or directory '%s' found", path)
	}
	return nil
}

func (cfg *Config) parseDest(path string) error {
	newPath := filepath.Join(cfg.HomeDir, path)

	var dirPath string

	ok := filepath.IsAbs(newPath)
	if !ok {
		err := createPath(newPath)
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "create path directory with parents at %s", dirPath)
	}
	return nil
}

func createPath(path string) error {
	dirPath := filepath.Dir(path)

	err := os.MkdirAll(dirPath, 0750)
	if err != nil {
		return fmt.Errorf("issue creating destination path %s: %v", dirPath, err)
	}

	// TODO: add verbose here

	return nil
}

func handlerError(err error) {
	fmt.Fprintln(os.Stderr, err)
	fmt.Fprintln(os.Stdout, "Try 'mvx --help' for more info.")
	os.Exit(1)
}
