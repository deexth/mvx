package cli

import (
	"fmt"
	"os"
	"path/filepath"
)

func handlerDestination(path string) (Dest, error) {
	return Dest{
		Loc: CleanPath(path),
	}, nil
}

func cleanDest(path string) Dest {
	return Dest{
		Loc: CleanPath(path),
	}
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
