package ops

import (
	"fmt"
	"os"

	"github.com/deexth/mvx/internal/cli"
)

func renameFile(src, dest string) error {
	destPath, err := cli.GetAbsPath(dest)
	if err != nil {
		return err
	}

	err = os.Rename(src, destPath)
	if err != nil {
		return fmt.Errorf("could not rename file: %v", err)
	}

	fmt.Fprintln(os.Stdout, "File renamed")

	return nil
}
