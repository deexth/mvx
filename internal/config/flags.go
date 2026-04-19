package config

import (
	flag "github.com/spf13/pflag"
)

var b = flag.Bool("no-clobber", false, "Do not overwrite an existing file")
