package ops

import (
	"fmt"
	"io"
	"io/fs"
	"os"
)

func fallbackForEXDEV(src SRC, finalDST string) error {
	if src.IsDir {
		return copyDir(src.FullPath, finalDST, src.Perm.Perm())
	}
	return copyFile(src.FullPath, finalDST, src.Perm.Perm())
}

func copyFile(src, dst string, srcPerm fs.FileMode) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("issue opening source file: %v", err)
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, srcPerm)
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

func copyDir(src, dst string, srcPerm fs.FileMode) error {
	// Recursively copy from src to dst
	return nil
}
