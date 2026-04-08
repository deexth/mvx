package main

import (
	"fmt"
	"os"

	"github.com/deexth/mvx/internal/cli"
)

func (cfg *Config) handlerSource(path string) (cli.Source, error) {
	cleanSrc := cli.CleanPath(path)
	absPath, err := cli.GetAbsPath(cleanSrc)
	if err != nil {
		return cli.Source{}, err
	}

	srcInfo, err := os.Stat(absPath)
	if err != nil {
		return cli.Source{}, fmt.Errorf("unable to get path info '%s': %v", cleanSrc, err)
	}

	return cli.Source{
		Loc:   cleanSrc,
		Name:  srcInfo.Name(),
		IsDir: srcInfo.IsDir(),
	}, nil
}
