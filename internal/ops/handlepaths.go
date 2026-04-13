package ops

import (
	"fmt"
	iofs "io/fs"
	"path/filepath"
	"strings"

	"github.com/deexth/mvx/internal/config"
	"github.com/deexth/mvx/internal/fs"
)

type SRC struct {
	FullPath string
	Name     string
	Mode     iofs.FileMode
	isDir    bool
}

type DST struct {
	FullPath string
	Name     string
	Mode     iofs.FileMode
	Exists   bool
	isDir    bool
}

func handlerSource(src []string, fs fs.FS) ([]SRC, error) {
	var srcInfos []SRC
	for _, source := range src {
		fullSrcPath, err := fs.Abs(source)
		if err != nil {
			return []SRC{}, fmt.Errorf("mvx: cannot move '%s': %v", source, err)
		}

		srcInfo, err := fs.Stat(fullSrcPath)
		if err != nil {
			return []SRC{}, fmt.Errorf("mv: cannot move '%s': %v", source, err)
		}

		srcInfos = append(srcInfos, SRC{
			FullPath: fullSrcPath,
			Name:     srcInfo.Name(),
			Mode:     srcInfo.Mode(),
			isDir:    srcInfo.IsDir(),
		})
	}

	return srcInfos, nil
}

func handlerDestination(dst string, fs fs.FS, cfg *config.Config) (DST, error) {
	newDST := expandPath(dst, cfg.HomeDir)

	_, err := fs.Abs(dstDir)
	if err != nil {
		return DST{}, fmt.Errorf("mvx: cannot move to '%s': %v", dst, err)
	}

	fullPath := filepath.Join(cfg.HomeDir, newDST)
	dstInfo, err := fs.Stat(fullPath)
	if err != nil {
		return DST{}, fmt.Errorf("mvx: cannot move to '%s': %v", newDST, err)
	}

	return DST{
		FullPath: fullPath,
		Name:     dstInfo.Name(),
		Mode:     dstInfo.Mode(),
		Exists:   true,
		isDir:    dstInfo.IsDir(),
	}, nil

}

func expandPath(path string, home string) string {
	ok := strings.HasPrefix(path, "~")
	if !ok {
		return filepath.Join(home, path)
	}

	newpath := strings.Split(path, "~")
	return filepath.Join(home, newpath[1])
}
