package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	src, dest, err := parseArgs(args[1:])
	if err != nil {
		handlerError(err)
	}

	fmt.Fprintf(os.Stdout, "source provide is %s and the destination is %s\n", src, dest)
}
