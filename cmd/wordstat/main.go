// Command wordstat prints the ten most frequent words from standard input.
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pikespeak/ci-patient/wordstat"
)

func main() { os.Exit(run(os.Stdin, os.Stdout, os.Stderr)) }

func run(in io.Reader, out, errOut io.Writer) int {
	top, err := wordstat.Top(in, 10)
	if err != nil {
		fmt.Fprintln(errOut, "wordstat:", err)
		return 1
	}
	for _, e := range top {
		fmt.Fprintf(out, "%d\t%s\n", e.Count, e.Word)
	}
	return 0
}
