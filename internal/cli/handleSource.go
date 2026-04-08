package cli

import (
	"fmt"
	"os"
)

func handlerSource(path string) (Source, error) {
	cleanSrc := CleanPath(path)
	absPath, err := GetAbsPath(cleanSrc)
	if err != nil {
		return Source{}, err
	}

	srcInfo, err := os.Stat(absPath)
	if err != nil {
		return Source{}, fmt.Errorf("unable to get path info '%s': %v", cleanSrc, err)
	}

	return Source{
		Loc:   cleanSrc,
		Name:  srcInfo.Name(),
		IsDir: srcInfo.IsDir(),
	}, nil
}
