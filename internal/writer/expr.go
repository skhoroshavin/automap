package writer

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/core/ast"
	"io"
)

func writeExpr(out io.Writer, expr ast.Expr, indent int) (err error) {
	switch x := expr.(type) {
	case *ast.ValueExpr:
		_, err = fmt.Fprint(out, x.Value)
		return
	case *ast.StructExpr:
		err = writeStructExpr(out, x, indent)
		return
	default:
		err = fmt.Errorf("unknown expression type %T", expr)
		return
	}
}

func writeStructExpr(out io.Writer, strct *ast.StructExpr, indent int) (err error) {
	if strct.IsPointer {
		_, err = fmt.Fprint(out, "&")
		if err != nil {
			return
		}
	}

	_, err = fmt.Fprintf(out, "%s{\n", strct.Name)
	if err != nil {
		return
	}

	for _, field := range strct.Fields {
		err = writeIndent(out, indent+1)
		if err != nil {
			return
		}

		_, err = fmt.Fprintf(out, "%s: ", field.Name)
		if err != nil {
			return
		}

		err = writeExpr(out, field.Value, indent+1)
		if err != nil {
			return
		}

		_, err = fmt.Fprintln(out, ",")
		if err != nil {
			return
		}
	}

	err = writeIndent(out, indent)
	if err != nil {
		return
	}
	_, err = fmt.Fprint(out, "}")
	return
}

func writeIndent(out io.Writer, indent int) (err error) {
	for i := 0; i != indent; i++ {
		if _, err = fmt.Fprint(out, "\t"); err != nil {
			return err
		}
	}
	return
}
