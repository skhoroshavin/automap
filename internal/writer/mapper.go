package writer

import (
	"errors"
	"fmt"
	"github.com/skhoroshavin/automap/internal/core/ast"
	"io"
)

func writeMapper(out io.Writer, mapper *ast.Mapper) (err error) {
	if mapper.Result == nil {
		return errors.New("mapper has empty return statement")
	}

	_, err = fmt.Fprintf(out, "\n%s {\n", mapper.Signature)

	for _, v := range mapper.Vars {
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

	err = writeExpr(out, mapper.Result, 1)
	if err != nil {
		return
	}

	_, err = fmt.Fprintln(out, "\n}")
	return
}
