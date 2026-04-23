package cli

import (
	"bufio"
	"fmt"
	"os"
)

func (o *MoveOptions) Interact(dst string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Fprintf(os.Stdout, "mvx: overwrite '%s'? ", dst)
	scanner.Scan()
	return scanner.Text()
}
