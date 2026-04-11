package main

import (
	"fmt"
	"os"

	"github.com/deexth/mvx/internal/config"
	"github.com/deexth/mvx/internal/fs"
	"github.com/deexth/mvx/internal/ops"
)

func main() {
	args := os.Args

	cfg, err := config.NewConfig(args[1:])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	err = ops.Move(&cfg, ops.MoveOptions{}, fs.OSFS{})
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("source: %s\ndestination: %s", cfg.Source[0:], cfg.Destination)
}
