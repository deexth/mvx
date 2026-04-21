// Package ops contains all the operations of the app
package ops

import (
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/deexth/mvx/internal/cli"
	"github.com/deexth/mvx/internal/config"
	"github.com/deexth/mvx/internal/fs"
)

func Move(cfg *config.Config, opts cli.MoveOptions, fs fs.FS) error {
	//
	srcs, err := HandlerSource(cfg.Source, fs)
	if err != nil {
		return err
	}

	dst, err := HandlerDestination(cfg.Destination, cfg.CWD, fs)
	if err != nil {
		return err
	}

	if len(srcs) > 1 && !dst.IsDir {
		return errors.New("mvx: target not a directory")
	}

	for _, src := range srcs {

		finalDst := ResolveDestination(
			src,
			dst,
		)

		err = fs.Rename(src.FullPath, finalDst)
		if err != nil {
			var linkErr *os.LinkError

			if errors.As(err, &linkErr) && linkErr.Err == syscall.EXDEV {
				// return fallbackForEXDEV(src, dst)
				// return errors.New("cross-device move detected, fallback not yet implemented")
				err = fallbackForEXDEV(src, finalDst)
				if err != nil {
					return err
				}

				_, err = fs.Stat(finalDst)
				if err != nil {
					return fmt.Errorf("file '%s' not moved to '%s'", src.FullPath, finalDst)
				}

				fs.Remove(src.FullPath)

				continue
			}

			return err
		}
	}

	return nil
}
