package main

import (
	"fmt"
	"os"
)

type Source struct {
	loc   string
	name  string
	isDir bool
}

func (cfg *Config) handlerSource(path string) (Source, error) {
	cleanSrc := cleanPath(path)
	absPath, err := getAbsPath(cleanSrc)
	if err != nil {
		return Source{}, err
	}

	srcInfo, err := os.Stat(absPath)
	if err != nil {
		return Source{}, fmt.Errorf("unable to get path info '%s': %v", cleanSrc, err)
	}

	return Source{
		loc:   cleanSrc,
		name:  srcInfo.Name(),
		isDir: srcInfo.IsDir(),
	}, nil
}
