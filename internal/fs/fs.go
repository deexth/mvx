// Package fs is a miner wrapper arround os
package fs

import "os"

type FS interface {
	Stat(name string) (os.FileInfo, error)
	Rename(oldName, newName string) error
	MkdirAll(path string, perm os.FileMode) error
	Remove(name string) error
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
