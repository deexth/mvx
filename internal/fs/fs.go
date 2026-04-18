// Package fs is a miner wrapper arround os
package fs

import (
	"io"
	"os"
	"path/filepath"
)

type FS interface {
	Stat(string) (os.FileInfo, error)
	Lstat(string) (os.FileInfo, error)
	Rename(oldPath, newPath string) error
	MkdirAll(path string, perm os.FileMode) error
	Remove(string) error
	Copy(dst io.Writer, src io.Reader) error
	Abs(string) (string, error)
	Dir(string) string
}

type OSFS struct{}

func (OSFS) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (OSFS) Lstat(name string) (os.FileInfo, error) {
	return os.Lstat(name)
}

func (OSFS) Rename(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
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
