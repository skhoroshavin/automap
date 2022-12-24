package writer

import (
	"automap/internal/utils"
	"io"
)

func writeImports(out io.StringWriter, imports []string) (err error) {
	if len(imports) == 0 {
		return
	}

	err = utils.WriteLn(out, "\nimport (")
	if err != nil {
		return
	}

	for _, s := range imports {
		err = utils.WriteLn(out, "\t%s", s)
		if err != nil {
			return
		}
	}

	err = utils.WriteLn(out, ")")
	return
}
