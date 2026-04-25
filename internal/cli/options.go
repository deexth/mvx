package cli

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func (o *MoveOptions) Interact(dst string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Fprintf(os.Stdout, "mvx: overwrite '%s'? ", dst)
	scanner.Scan()
	return scanner.Text()
}

func (o *MoveOptions) Upd(src, dst time.Time) bool {
	return src.After(dst)
}
