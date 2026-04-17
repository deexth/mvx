// Package config; contains the core configuration of the app
package config

import (
	"errors"
	"os"
)

type Config struct {
	CWD         string
	Source      []string
	Destination string
}

func NewConfig(args []string) (Config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}

	if len(args) < 2 {
		return Config{}, errors.New("usage: mvx <source>... <destination> ")
	}

	if len(args) > 2 {
		return Config{
			CWD:         cwd,
			Source:      args[0 : len(args)-1],
			Destination: args[len(args)-1],
		}, nil
	}

	return Config{
		CWD:         cwd,
		Source:      []string{args[0]},
		Destination: args[1],
	}, nil
}
