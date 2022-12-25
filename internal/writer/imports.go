package writer

import (
	"fmt"
	"io"
)

func writeImports(out io.Writer, imports []string) (err error) {
	if len(imports) == 0 {
		return
	}

	_, err = fmt.Fprintln(out, "\nimport (")
	if err != nil {
		return
	}

	for _, s := range imports {
		_, err = fmt.Fprintf(out, "\t%s\n", s)
		if err != nil {
			return
		}
	}

	_, err = fmt.Fprintln(out, ")")
	return
}
