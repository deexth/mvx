// Package cli contains implementation of all cli interactivity, e.g: flags
package cli

import flag "github.com/spf13/pflag"

type MoveOptions struct {
	NoClobber   bool
	Interactive bool
	Verbose     bool
	Update      bool
	Force       bool
	Tree        bool
	Preview     bool
	Diff        bool
	Help        bool
	Copy        bool
	Parents     bool
	Backup      bool
}

func InitFlags() MoveOptions {
	noClobber := flag.BoolP("no-clobber", "n", false, "Do not overwrite an existing file")
	interactive := flag.BoolP("interactive", "i", false, "Prompt before overwrite")
	force := flag.BoolP("force", "f", false, "Do not prompt before overwriting")
	verbose := flag.BoolP("verbose", "v", false, "Explain what is being done")
	update := flag.BoolP("update", "u", false, "Move only when the SOURCE file is newer than the destination file or when the destination file is missing")
	tree := flag.BoolP("tree", "t", false, "Display the destination tree after move")
	preview := flag.BoolP("preview", "p", false, "Display the destination file after move. Can be used only if 1 file is moved")
	diff := flag.BoolP("diff", "d", false, "Displays the difference between the source and target files after a move")
	copy := flag.BoolP("copy", "c", false, "Used to keep a copy of the source")
	parents := flag.BoolP("parents", "P", false, "Create the parent dirs of the destination if not exist")
	backup := flag.BoolP("backup", "b", false, "Create backup of the source. Used only on files")
	help := flag.BoolP("help", "h", false, "Display manual for mvx")

	flag.Parse()

	return MoveOptions{
		NoClobber:   *noClobber,
		Interactive: *interactive,
		Verbose:     *verbose,
		Force:       *force,
		Update:      *update,
		Tree:        *tree,
		Preview:     *preview,
		Parents:     *parents,
		Diff:        *diff,
		Backup:      *backup,
		Copy:        *copy,
		Help:        *help,
	}
}
