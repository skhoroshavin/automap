package writer

import (
	"fmt"
	"io"
)

func writeOldHeader(out io.Writer, packageName string) (err error) {
	_, err = fmt.Fprintf(out, packageHeader, packageName)
	return
}
