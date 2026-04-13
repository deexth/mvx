// Package pathutil contains the logic for path resolving before moving/renaming
package pathutil

import (
	"path/filepath"
)

func ResolveDestination(src, dst string, dstExists, dstIsDir bool) string {
	if dstExists && dstIsDir {
		return filepath.Join(dst, src)
	}

	return dst
}

func ValidateDestination(dstPath string) (string, error) {
	return "", nil
}
