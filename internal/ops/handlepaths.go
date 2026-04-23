package ops

import (
	"fmt"
	iofs "io/fs"
	"os"
	"time"

	"github.com/deexth/mvx/internal/fs"
)

type SRC struct {
	ModTime time.Time
	Perm    iofs.FileMode
	Path
}

type DST struct {
	Path
}

type Path struct {
	FullPath string
	Name     string
	IsDir    bool
	Exists   bool
}

func HandlerSource(src []string, fs fs.FS) ([]SRC, error) {
	var srcInfos []SRC
	for _, source := range src {
		fullSrcPath, err := fs.Abs(source)
		if err != nil {
			return []SRC{}, fmt.Errorf("mvx: cannot move '%s': %v", source, err)
		}

		srcInfo, err := fs.Lstat(fullSrcPath)
		if err != nil {
			return []SRC{}, fmt.Errorf("mv: cannot move '%s': %v", source, err)
		}

		srcInfos = append(srcInfos, SRC{
			ModTime: srcInfo.ModTime(),
			Perm:    srcInfo.Mode().Perm(),
			Path: Path{
				FullPath: fullSrcPath,
				Name:     srcInfo.Name(),
				IsDir:    srcInfo.IsDir(),
				Exists:   true,
			},
		})
	}

	return srcInfos, nil
}

func HandlerDestination(dst, cwd string, fs fs.FS) (DST, error) {
	path, err := expandPath(dst, cwd)
	if err != nil {
		return DST{}, fmt.Errorf("mvx: cannot move to '%s': %v", dst, err)
	}

	dstInfo, err := fs.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return validateDestination(path, dst, fs, false)
		}
		return DST{}, fmt.Errorf("mvx: cannot move to '%s': %v", path, err)
	}

	return DST{
		Path: Path{
			FullPath: path,
			Name:     dst,
			IsDir:    dstInfo.IsDir(),
			Exists:   true,
		}}, nil

}
