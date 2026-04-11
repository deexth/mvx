// Package config; contains the core configuration of the app
package config

import (
	"errors"
)

type Config struct {
	Source      []string
	Destination string
}

func NewConfig(args []string) (Config, error) {
	if len(args) < 2 {
		return Config{}, errors.New("usage: mvx <source>... <destination> ")
	}

	if len(args) > 2 {
		return Config{
			Source:      args[0 : len(args)-1],
			Destination: args[len(args)-1],
		}, nil
	}

	return Config{
		Source:      []string{args[0]},
		Destination: args[1],
	}, nil
}
