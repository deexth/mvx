// Package cli contains implementation of all cli interactivity, e.g: flags
package cli

import flag "github.com/spf13/pflag"

type MoveOptions struct {
	NoClobber       bool
	Interactive     bool
	Verbose         bool
	Update          bool
	Force           bool
	Backup          bool
	NoTargetDirect  bool
	Version         bool
	TargetDirectory bool
	Suffix          bool
	Help            bool
	Context         bool
	// mvx added flags
	Tree    bool
	Preview bool
	Diff    bool
	Copy    bool
	Parents bool
}

func InitFlags() MoveOptions {
	noClobber := flag.BoolP("no-clobber", "n", false, "do not overwrite an existing file")
	interactive := flag.BoolP("interactive", "i", false, "prompt before overwrite")
	force := flag.BoolP("force", "f", false, "do not prompt before overwriting")
	verbose := flag.BoolP("verbose", "v", false, "explain what is being done")
	update := flag.BoolP("update", "u", false, "move only when the SOURCE file is newer than the destination file or when the destination file is missing")
	tree := flag.BoolP("tree", "t", false, "display the destination tree after move")
	preview := flag.BoolP("preview", "p", false, "display the destination file after move. Can be used only if 1 file is moved")
	diff := flag.BoolP("diff", "d", false, "displays the difference between the source and target files after a move")
	copy := flag.BoolP("copy", "c", false, "used to keep a copy of the source")
	parents := flag.BoolP("parents", "P", false, "create the parent dirs of the destination if not existing")
	backup := flag.BoolP("backup", "b", false, "make a backup of each existing destination file")
	help := flag.BoolP("help", "h", false, "display usage information for mvx")
	ntd := flag.BoolP("no-target-directory", "T", false, "treat final argument as a normal file")
	version := flag.BoolP("version", "V", false, "treat final argument as a normal file")
	context := flag.BoolP("context", "z", false, "set SELinux security context of destination file to default type")
	td := flag.BoolP("target-directory", "D", false, "move all source arguments into specified directory")
	suffix := flag.BoolP("suffix", "x", false, "specify the backup suffix")

	flag.Parse()

	return MoveOptions{
		NoClobber:       *noClobber,
		Interactive:     *interactive,
		Verbose:         *verbose,
		Force:           *force,
		Update:          *update,
		Tree:            *tree,
		Preview:         *preview,
		Parents:         *parents,
		Diff:            *diff,
		Backup:          *backup,
		Copy:            *copy,
		Help:            *help,
		NoTargetDirect:  *ntd,
		Version:         *version,
		Context:         *context,
		TargetDirectory: *td,
		Suffix:          *suffix,
	}
}
