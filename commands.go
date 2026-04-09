package main

import (
	"github.com/deexth/mvx/internal/cli"
	"github.com/deexth/mvx/internal/ops"
)

func actions(cfg *cli.Config) error {
	if len(cfg.PathArgs.Src) == 1 && !cfg.PathArgs.Src[0].IsDir {
		err := ops.RenameFile(cfg.PathArgs.Src[0], cfg.PathArgs.Dest)
		if err != nil {
			return err
		}
	}

	return nil
}
