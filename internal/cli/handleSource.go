package cli

import (
	"fmt"
	"os"
)

func handlerSource(path string) (Source, error) {
	absPath, err := GetAbsPath(path)
	if err != nil {
		return Source{}, err
	}

	srcInfo, err := os.Stat(absPath)
	if err != nil {
		return Source{}, fmt.Errorf("unable to get path info '%s': %v", path, err)
	}

	return Source{
		Loc:   absPath,
		Name:  srcInfo.Name(),
		IsDir: srcInfo.IsDir(),
	}, nil
}
