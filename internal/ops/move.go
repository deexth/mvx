// Package ops contains all the operations of the app
package ops

import (
	"github.com/deexth/mvx/internal/cli"
	"github.com/deexth/mvx/internal/config"
	"github.com/deexth/mvx/internal/fs"
)

func Move(cfg *config.Config, opts cli.MoveOptions, fs fs.FS) error {
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
	err := handlerMove(cfg, fs)
	if err != nil {
		return err
	}
	return nil
}
