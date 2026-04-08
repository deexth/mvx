package main

import (
	"github.com/deexth/mvx/internal/cli"
	"github.com/deexth/mvx/internal/views"
	"github.com/spf13/pflag"
)

type Flag struct {
	name        string
	description string
	callback    func(string) error
}

type Flags struct {
	f map[string]Flag
}

func initFlags(cfg *cli.Config) {
	laodedFlags := loadFlags(cfg)

	for _, f := range laodedFlags.f {
		pflag.Func(f.name, f.description, f.callback)
	}

	pflag.Parse()
}

func loadFlags(cfg *cli.Config) Flags {
	return Flags{
		f: map[string]Flag{
			"tree": Flag{
				name:        "tree",
				description: "--tree recursively lists the destination directory",
				callback:    middlewareFlags(cfg.HomeDir, views.Tree),
			},
		},
	}
}

func middlewareFlags(path string, fn func(string) error) func(string) error {
	return func(s string) error {
		return fn(path)
	}
}
