package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Dest struct {
	loc string
}

func (cfg *Config) handlerDestination(path string) (Dest, error) {
	return Dest{
		loc: cleanPath(path),
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
