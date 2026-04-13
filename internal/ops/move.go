// Package ops contains all the operations of the app
package ops

import (
	"github.com/deexth/mvx/internal/config"
	"github.com/deexth/mvx/internal/fs"
	pu "github.com/deexth/mvx/internal/pathutil"
)

type MoveOptions struct {
	Copy    bool
	Force   bool
	Parents bool
}

func Move(cfg *config.Config, ops MoveOptions, fs fs.FS) error {
	// switch ops {
	// case MoveOptions{
	// 	true,
	// 	false,
	// 	false,
	// }:
	// 	return nil
	// case MoveOptions{
	// 	false,
	// 	true,
	// 	false,
	// }:
	// 	return nil
	// case MoveOptions{
	// 	false,
	// 	false,
	// 	true,
	// }:
	// 	return nil
	// default:
	// 	return nil
	// }

	srcs, err := handlerSource(cfg.Source, fs)
	if err != nil {
		return err
	}

	dst, err := handlerDestination(cfg.Destination, cfg.HomeDir, fs)
	if err != nil {
		return err
	}

	for _, src := range srcs {
		finalDst := pu.ResolveDestination(
			src.Info.Name(),
			dst.FullPath,
			dst.Exists,
			dst.Info.IsDir(),
		)

	}

	return nil
}
