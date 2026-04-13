// Package fs is a miner wrapper arround os
package fs

import (
	"io"
	"os"
	"path/filepath"
)

type FS interface {
	Stat(name string) (os.FileInfo, error)
	Rename(oldName, newName string) error
	MkdirAll(path string, perm os.FileMode) error
	Remove(name string) error
	Copy(dst io.Writer, src io.Reader) error
	Abs(path string) (string, error)
	Dir(string) string
}

type OSFS struct{}

func (OSFS) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (OSFS) Rename(oldName, newName string) error {
	return os.Rename(oldName, newName)
}

func (OSFS) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (OSFS) Remove(name string) error {
	return os.Remove(name)
}

func (OSFS) Copy(dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func (OSFS) Abs(path string) (string, error) {
	return filepath.Abs(path)
}

func (OSFS) Dir(path string) string {
	return filepath.Dir(path)
}
