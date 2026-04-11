// Package ops contains all the operations of the app
package ops

import (
	"github.com/deexth/mvx/internal/config"
	"github.com/deexth/mvx/internal/fs"
)

type MoveOptions struct {
	Copy    bool
	Force   bool
	Parents bool
}

func Move(cfg *config.Config, ops MoveOptions, fs fs.FS) error {
	if len(cfg.Source) > 1 {

	}
	return nil
}
