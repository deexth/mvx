package ops

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/deexth/mvx/internal/fs"
)

func ResolveDestination(src SRC, dst DST) string {
	if dst.IsDir {
		path := filepath.Join(dst.FullPath, src.Name)
		return path
	}

	return dst.FullPath
}

func validateDestination(dstPath string, fs fs.FS, createParent bool) (DST, error) {
	dstDIR := fs.Dir(dstPath)
	_, err := fs.Stat(dstDIR)
	if err != nil && !createParent {
		return DST{}, fmt.Errorf("mvx: cannot move to '%s': %v", dstPath, err)
	}

	return DST{
		Path: Path{FullPath: dstPath,
			Exists: false,
			IsDir:  false,
		}}, nil

}

func expandPath(path string, cwd string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if path == "~" || strings.HasPrefix(path, "~/") {
		return strings.Replace(path, "~", home, 1), nil
	}

	if strings.HasPrefix(path, home) {
		return path, nil
	}

	return filepath.Join(cwd, path), nil
}
