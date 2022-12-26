package writer

import (
	"errors"
	"fmt"
	"github.com/skhoroshavin/automap/internal/core/ast"
	"io"
)

func writeFunc(out io.Writer, fn *ast.Mapper) (err error) {
	if fn.Result == nil {
		return errors.New("function has empty return statement")
	}

	for _, v := range fn.Vars {
		_, err = fmt.Fprintf(out, "\t%s := ", v.Name)
		if err != nil {
			return
		}

		err = writeExpr(out, v.Value, 1)
		if err != nil {
			return
		}

		_, err = fmt.Fprintln(out)
		if err != nil {
			return
		}
	}

	_, err = fmt.Fprint(out, "\treturn ")
	if err != nil {
		return
	}

	err = writeExpr(out, fn.Result, 1)
	if err != nil {
		return
	}

	_, err = fmt.Fprintln(out)
	return
}
