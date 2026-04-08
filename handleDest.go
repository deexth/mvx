package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/deexth/mvx/internal/cli"
)

func (cfg *Config) handlerDestination(path string) (cli.Dest, error) {
	return cli.Dest{
		Loc: cli.CleanPath(path),
	}, nil
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
