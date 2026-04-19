// Package ops contains all the operations of the app
package ops

import (
	"errors"
	"fmt"
	"io"
	"os"
	"syscall"

	"github.com/deexth/mvx/internal/config"
	"github.com/deexth/mvx/internal/fs"
)

type MoveOptions struct {
	NoClobber   bool
	Interactive bool
	Verbose     bool
	Update      bool
	Force       bool
	Tree        bool
	Preview     bool
	Diff        bool
	Help        bool
	Copy        bool
	Parents     bool
}

func Move(cfg *config.Config, opts MoveOptions, fs fs.FS) error {
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

func fallbackForEXDEV(src SRC, finalDST string) error {
	srcFile, err := os.Open(src.FullPath)
	if err != nil {
		return fmt.Errorf("issue opening source file: %v", err)
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(finalDST, os.O_RDWR|os.O_CREATE, src.Perm)
	if err != nil {
		return fmt.Errorf("issue creating destination file: %v", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("issue copying content from source file: %v", err)
	}

	return nil
}
