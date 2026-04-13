package ops

import (
	"fmt"
	iofs "io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/deexth/mvx/internal/config"
	"github.com/deexth/mvx/internal/fs"
)

type Path struct {
	FullPath string
	Info     iofs.FileInfo
	Exists   bool
}

func handlerSource(src []string, fs fs.FS) ([]Path, error) {
	var srcInfos []Path
	for _, source := range src {
		fullSrcPath, err := fs.Abs(source)
		if err != nil {
			return []Path{}, fmt.Errorf("mvx: cannot move '%s': %v", source, err)
		}

		srcInfo, err := fs.Stat(fullSrcPath)
		if err != nil {
			return []Path{}, fmt.Errorf("mv: cannot move '%s': %v", source, err)
		}

		srcInfos = append(srcInfos, Path{
			FullPath: fullSrcPath,
			Info:     srcInfo,
			Exists:   true,
		})
	}

	return srcInfos, nil
}

func handlerDestination(dst string, fs fs.FS, cfg *config.Config) (Path, error) {
	path := expandPath(dst, cfg.HomeDir)

	dstInfo, err := fs.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Path{
				FullPath: path,
				Exists:   false,
				Info:     nil,
			}, nil
		}
		return Path{}, fmt.Errorf("mvx: cannot move to '%s': %v", path, err)
	}

	return Path{
		FullPath: path,
		Info:     dstInfo,
		Exists:   true,
	}, nil

}

func expandPath(path string, home string) string {
	if path == "~" {
		return home
	}

	newPath, _ := strings.CutPrefix(path, "~/")

	return filepath.Join(home, newPath)
}
