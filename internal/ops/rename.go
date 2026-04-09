package ops

import (
	"fmt"
	"os"

	"github.com/deexth/mvx/internal/cli"
)

func RenameFile(src cli.Source, dest cli.Dest) error {
	destPath, err := cli.GetAbsPath(dest.Loc)
	if err != nil {
		return err
	}

	err = os.Rename(src.Loc, destPath)
	if err != nil {
		return fmt.Errorf("could not rename file: %v", err)
	}

	fmt.Fprintln(os.Stdout, "File renamed")

	return nil
}
