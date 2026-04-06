package main

import (
	"fmt"
	"os"
)

func (cfg *Config) renameFile(src, dest string) error {
	destPath, err := getAbsPath(dest)
	if err != nil {
		return err
	}

	err = os.Rename(src, destPath)
	if err != nil {
		return fmt.Errorf("could not rename file: %v", err)
	}

	fmt.Fprintln(os.Stdout, "File renamed")

	return nil
}
