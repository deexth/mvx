package ops

import (
	"fmt"
	"io"
	"os"
)

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
