package ops

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/deexth/mvx/internal/fs"
)

func ResolveDestination(src SRC, dst DST) string {
	if dst.Exists && dst.IsDir {
		path := filepath.Join(dst.FullPath, src.Name)
		fmt.Fprintf(os.Stdout, "mvx: moving to '%s'", path)
		return path
	}

	return dst.FullPath
}

func validateDestination(dstPath string, fs fs.FS, createParent bool) (DST, error) {
	dstDIR := fs.Dir(dstPath)
	_, err := fs.Abs(dstDIR)
	if err != nil && !createParent {
		return DST{}, fmt.Errorf("mvx: cannot move to '%s': %v", dstPath, err)
	}

	return DST{
		Path: Path{FullPath: dstPath,
			Exists: true,
			IsDir:  false,
		}}, nil

}

func expandPath(path string, home string) string {
	if path == "~" {
		return home
	}

	newPath, _ := strings.CutPrefix(path, "~/")

	return filepath.Join(home, newPath)
}
