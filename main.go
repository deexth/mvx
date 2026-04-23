package main

import (
	"fmt"
	"os"

	"github.com/deexth/mvx/internal/cli"
	"github.com/deexth/mvx/internal/config"
	"github.com/deexth/mvx/internal/fs"
	"github.com/deexth/mvx/internal/ops"
	flag "github.com/spf13/pflag"
)

func main() {
	opts := cli.InitFlags()

	args := flag.Args()

	cfg, err := config.NewConfig(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	err = ops.Move(&cfg, opts, fs.OSFS{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
