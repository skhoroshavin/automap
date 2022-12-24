package utils

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"io"
)

func AST2String(expr ast.Node) string {
	buf := bytes.Buffer{}
	fset := token.NewFileSet()
	err := printer.Fprint(&buf, fset, expr)
	if err != nil {
		return ""
	}
	return buf.String()
}

func WriteLn(out io.StringWriter, format string, args ...interface{}) (err error) {
	_, err = out.WriteString(fmt.Sprintf(format, args...))
	if err != nil {
		return err
	}
	_, err = out.WriteString("\n")
	return
}
